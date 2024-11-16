package tailless

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	typeVariable    = 1
	typeAtRule      = 2
	typeDeclaration = 3
	typeSelector    = 4
	typeOpenBrace   = 5
	typeCloseBrace  = 6
	typeImport      = 7
	typeMixin       = 8
)

var reVariable = regexp.MustCompile(`@[0-9A-Za-z-_]+`)
var reBraces = regexp.MustCompile(`[\{\}]`)

type parser struct {
	Elements *[]element
}

type line struct {
	Text       string
	LineNumber int
}

type lines struct {
	Items []line
}

func newLines() *lines {
	lines := lines{}
	lines.Items = make([]line, 0)
	return &lines
}

type element struct {
	Text        string
	ElementType int
	LineNumber  int
	Level       int
}

type elements struct {
	Items []element
}

func newElements() *elements {
	elements := elements{}
	elements.Items = make([]element, 0)
	return &elements
}

func newParser() *parser {
	return &parser{}
}

func (p *parser) Parse(srcFilename string, destFilename string) error {
	lines, err := p.RemoveComments(srcFilename)
	if err != nil {
		return err
	}

	lines, err = p.SplitBraces(lines)
	if err != nil {
		return err
	}

	elements, err := p.SplitIntoElements(lines)
	if err != nil {
		return err
	}

	err = p.ValidateElements(elements)
	if err != nil {
		return err
	}

	tree, err := p.BuildTree(elements)
	if err != nil {
		return err
	}

	err = resolveMixins(tree)
	if err != nil {
		return err
	}

	err = resolveVariables(tree)
	if err != nil {
		return err
	}

	expandSelectors(tree)

	tree.HideIfEmpty()

	file, err := os.Create(destFilename)
	if err != nil {
		return err
	}

	defer file.Close()

	tree.Render(file)

	file.Close()

	return nil
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func (p *parser) RemoveComments(filename string) (*lines, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := newLines()
	lineNumber := 0
	insideComment := false

	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		if !insideComment && line != "" && line[:1] == "@" {
			lines.Add(line, lineNumber)
			continue
		}

		pos := strings.Index(line, "//")
		if pos >= 0 {
			line = line[0:pos]
		}

		posStart := strings.Index(line, "/*")
		posEnd := strings.LastIndex(line, "*/")

		if insideComment {
			if posEnd >= 0 {
				insideComment = false
				lines.Add(line[posEnd+2:], lineNumber)
			} else {
				continue
			}
		} else {
			if posStart >= 0 && posEnd >= 0 {
				before := line[0:posStart]
				after := line[posEnd+2:]

				lines.Add(before, lineNumber)
				lines.Add(after, lineNumber)
			} else if posStart >= 0 {
				insideComment = true
				lines.Add(line[0:posStart], lineNumber)
			} else {
				lines.Add(line, lineNumber)
			}
		}
	}

	return lines, nil
}

func (p *parser) SplitBraces(lines *lines) (*lines, error) {
	newLines := newLines()

	for _, line := range lines.Items {
		text := line.Text
		if text == "{" || text == "}" {
			newLines.AddLine(line)
			continue
		}

		matches := reBraces.FindAllStringIndex(text, -1)

		if len(matches) == 0 {
			newLines.AddLine(line)
			continue
		}

		var flat []int
		for _, inner := range matches {
			flat = append(flat, inner...)
		}

		start := 0
		for _, pos := range flat {
			newLines.Add(text[start:pos], line.LineNumber)
			start = pos
		}

		newLines.Add(text[start:], line.LineNumber)
	}

	return newLines, nil

}

func (p *parser) SplitIntoElements(lines *lines) (*elements, error) {
	elements := newElements()

	inDeclaration := false
	level := 0

	for _, line := range lines.Items {
		str := line.Text

		if isVariable(str) {
			elements.Add(str, typeVariable, line.LineNumber)
		} else if isAtRule(str) {
			if str[:7] == "@import" {
				elements.Add(str, typeImport, line.LineNumber)
			} else {
				elements.Add(str, typeAtRule, line.LineNumber)
			}
		} else if isDeclarationStart(str) {
			elements.Add(str, typeDeclaration, line.LineNumber)
			inDeclaration = !endsWithSemiColon(str)
		} else if str == "{" {
			level++
			elements.AddLevel(str, typeOpenBrace, line.LineNumber, level)
		} else if str == "}" {
			elements.AddLevel(str, typeCloseBrace, line.LineNumber, level)
			level--
		} else if inDeclaration {
			elements.Add(str, typeDeclaration, line.LineNumber)
			inDeclaration = !endsWithSemiColon(str)
		} else if endsWithSemiColon(str) {
			elements.Add(str, typeMixin, line.LineNumber)
		} else {
			elements.Add(str, typeSelector, line.LineNumber)
		}
	}

	return elements, nil
}

func (p *parser) ValidateElements(elements *elements) error {
	lastLevel := 1
	for _, element := range elements.Items {
		if element.ElementType != typeCloseBrace {
			continue
		}

		level := element.Level
		if level <= 0 {
			return fmt.Errorf("Line %d: Unexpected brace", element.LineNumber)
		}

		lastLevel = level
	}

	items := elements.Items
	count := len(items)
	for i := 0; i < count-1; i++ {
		item := items[i]
		nextItem := items[i+1]

		elementType := item.ElementType
		nextElementType := nextItem.ElementType

		if elementType == typeAtRule && nextElementType != typeOpenBrace {
			return fmt.Errorf("Line %d: Missing opening brace", nextItem.LineNumber)
		}

		if elementType == typeSelector {
			if nextElementType != typeSelector && nextElementType != typeOpenBrace {
				return fmt.Errorf("Line %d: Missing opening brace", nextItem.LineNumber)
			}
		}

		if nextElementType == typeOpenBrace {
			if elementType != typeAtRule && elementType != typeSelector {
				fmt.Println("Type: ", elementType)
				return fmt.Errorf("Line %d: Invalid opening brace", nextItem.LineNumber)
			}
		}
	}

	if lastLevel != 1 {
		return fmt.Errorf("Missing braces at end of file")
	}

	return nil
}

func isVariable(str string) bool {
	if str[0:1] != "@" {
		return false
	}

	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		return false
	}

	name := strings.TrimSpace(parts[0])
	if reVariable.FindString(name) != name {
		return false
	}

	return true
}

func isAtRule(str string) bool {
	if str[0:1] != "@" {
		return false
	}

	return true
}

func isDeclarationStart(str string) bool {
	pos := strings.Index(str, ":")
	if pos < 0 {
		return false
	}

	if pos == len(str)-1 {
		return true
	}

	if str[pos+1:pos+2] == " " {
		return true
	}

	return false
}

func endsWithSemiColon(str string) bool {
	if str[len(str)-1:] != ";" {
		return false
	}

	return true
}

func (l *lines) Add(text string, lineNumber int) {
	text = strings.TrimSpace(text)
	if text == "" {
		return
	}

	l.Items = append(l.Items, line{text, lineNumber})
}

func (l *lines) AddLine(line line) {
	l.Items = append(l.Items, line)
}

func (e *elements) Add(text string, elementType int, lineNumber int) {
	e.AddLevel(text, elementType, lineNumber, 0)
}

func (e *elements) AddLevel(text string, elementType int, lineNumber int, level int) {
	e.Items = append(e.Items, element{text, elementType, lineNumber, level})
}
