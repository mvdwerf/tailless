package tailless

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

type node interface {
	AddChild(node)
	GetChildren() []node
	SetChildren([]node)
	Dump(string)
	ExpandSelectors([]string)
	HideIfEmpty() bool
	Render(io.Writer)
	GetType() string
	GetVariable(*variablesCollection)
	ReplaceVariables(*variablesCollection) error
	GetMixin(mixins)
	GetMixinName() string
	GetCopy() node
	GetLineNumber() int
	IsParentOf(node) bool
}

type baseNode struct {
	Children   []node
	LineNumber int
	Hidden     bool
}

func (n *baseNode) ExpandSelectors(parentSelectors []string) {

}

func (n *baseNode) HideIfEmpty() bool {
	return false
}

func (n *baseNode) GetChildren() []node {
	return n.Children
}

func (n *baseNode) SetChildren(children []node) {
	n.Children = children
}

func (n *baseNode) AddChild(child node) {
	n.Children = append(n.Children, child)
}

func (n *baseNode) GetVariable(variables *variablesCollection) {

}

func (n *baseNode) ReplaceVariables(variables *variablesCollection) error {
	return nil
}

func (n *baseNode) GetMixin(mixins) {

}

func (n *baseNode) GetMixinName() string {
	return ""
}

func (n *baseNode) GetCopy() node {
	return nil
}

func (n *baseNode) GetLineNumber() int {
	return n.LineNumber
}

func (n *baseNode) Dump(indent string) {
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

func (n *baseNode) Render(io.Writer) {

}

func (n *baseNode) GetType() string {
	return ""
}

func (n *baseNode) IsParentOf(node node) bool {
	for _, child := range n.Children {
		if child == node {
			return true
		}

		if child.IsParentOf(node) {
			return true
		}
	}

	return false
}

type rootNode struct {
	baseNode
}

func expandSelectors(tree *rootNode) {
	for _, child := range tree.Children {
		child.ExpandSelectors([]string{""})
	}
}

func (n *rootNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *rootNode) Render(w io.Writer) {
	for _, child := range n.Children {
		child.Render(w)
	}
}

type variableNode struct {
	baseNode
	Text string
}

func (n *variableNode) GetType() string {
	return "variable"
}

func (n *variableNode) HideIfEmpty() bool {
	n.Hidden = true
	return true
}

func (n *variableNode) GetVariable(variables *variablesCollection) {
	parts := strings.Split(n.Text, ":")
	if len(parts) != 2 {
		return
	}

	name := strings.TrimPrefix(strings.TrimSpace(parts[0]), "@")
	value := strings.TrimSuffix(strings.TrimSpace(parts[1]), ";")

	variables.Set(name, value)
}

func (n *variableNode) GetCopy() node {
	return newVariableNode(n.Text, n.LineNumber)
}

func (n *variableNode) Dump(indent string) {
	fmt.Printf("%sVariableNode: %s\n", indent, n.Text)
}

type selectorNode struct {
	baseNode
	Selectors       []string
	MergedSelectors []string
}

func (n *selectorNode) ExpandSelectors(parentSelectors []string) {
	selectors := mergeSelectors(parentSelectors, n.Selectors)
	n.MergedSelectors = selectors
	for _, child := range n.Children {
		child.ExpandSelectors(selectors)
	}
}

func (n *selectorNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *selectorNode) GetMixin(mixins mixins) {
	selectors := n.Selectors
	if len(selectors) != 1 {
		return
	}

	name := strings.TrimSuffix(selectors[0], "()")

	mixins.Set(name, n)
}

func (n *selectorNode) GetCopy() node {
	copy := newSelectorNode(n.Selectors, n.LineNumber)

	for _, child := range n.Children {
		copy.Children = append(copy.Children, child.GetCopy())
	}

	return copy
}

func (n *selectorNode) Dump(indent string) {
	fmt.Printf("%sSelectorNode", indent)

	for _, selector := range n.MergedSelectors {
		fmt.Printf(" %s -", selector)
	}

	hidden := ""
	if n.Hidden {
		hidden = " (hidden)"
	}

	fmt.Println(hidden)

	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

func (n *selectorNode) Render(w io.Writer) {
	if n.Hidden {
		return
	}

	declarationNodes := make([]node, 0)
	for _, child := range n.Children {
		if child.GetType() == "declaration" {
			declarationNodes = append(declarationNodes, child)
		}
	}

	if len(declarationNodes) > 0 {

		last := len(n.MergedSelectors) - 1

		for i, selector := range n.MergedSelectors {
			if i == last {
				fmt.Fprintf(w, "%s {\n", selector)
			} else {
				fmt.Fprintf(w, "%s,\n", selector)
			}
		}

		for _, child := range declarationNodes {
			child.Render(w)
		}

		fmt.Fprintf(w, "}\n")
	}

	for _, child := range n.Children {
		if child.GetType() != "declaration" {
			child.Render(w)
		}
	}
}

type declarationNode struct {
	baseNode
	Text string
}

func (n *declarationNode) GetType() string {
	return "declaration"
}

func (n *declarationNode) GetCopy() node {
	return newDeclarationNode(n.Text, n.LineNumber)
}

func (n *declarationNode) Render(w io.Writer) {
	fmt.Fprintf(w, "  %s\n", n.Text)
}

func (n *declarationNode) ReplaceVariables(variables *variablesCollection) error {
	text := n.Text

	for {
		match := reVariable.FindStringIndex(text)
		if match == nil {
			n.Text = text
			return nil
		}

		start := match[0]
		end := match[1]

		name := text[start+1 : end]
		value := variables.Get(name)

		if value == "" {
			return fmt.Errorf("Line %d: Variable '%s' not found", n.LineNumber, name)
		}

		text = text[0:start] + value + text[end:]
	}
}

func (n *declarationNode) Dump(indent string) {
	fmt.Printf("%sDeclarationNode: %s\n", indent, n.Text)
}

type mixinNode struct {
	baseNode
	Text string
}

func (n *mixinNode) ExpandSelectors(parentSelectors []string) {
	for _, child := range n.Children {
		child.ExpandSelectors(parentSelectors)
	}
}

func (n *mixinNode) GetMixinName() string {
	return strings.TrimSuffix(strings.TrimSuffix(n.Text, ";"), "()")
}

func (n *mixinNode) Render(w io.Writer) {
	for _, child := range n.Children {
		child.Render(w)
	}
}

func (n *mixinNode) Dump(indent string) {
	fmt.Printf("%sMixinNode: %s\n", indent, n.Text)
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

type atRuleNode struct {
	baseNode
	Text            string
	ParentSelectors []string
}

func (n *atRuleNode) ExpandSelectors(parentSelectors []string) {
	n.ParentSelectors = parentSelectors
	for _, child := range n.Children {
		child.ExpandSelectors(parentSelectors)
	}
}

func (n *atRuleNode) ReplaceVariables(variables *variablesCollection) error {
	text := n.Text[1:]

	for {
		match := reVariable.FindStringIndex(text)
		if match == nil {
			n.Text = "@" + text
			return nil
		}

		start := match[0]
		end := match[1]

		name := text[start+1 : end]
		value := variables.Get(name)

		if value == "" {
			return fmt.Errorf("Line %d: Variable '%s' not found", n.LineNumber, name)
		}

		text = text[0:start] + value + text[end:]
	}
}

func (n *atRuleNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *atRuleNode) Render(w io.Writer) {
	fmt.Fprintf(w, "%s {\n", n.Text)

	declarationNodes := make([]node, 0)
	for _, child := range n.Children {
		if child.GetType() == "declaration" {
			declarationNodes = append(declarationNodes, child)
		}
	}

	if len(declarationNodes) > 0 {
		last := len(n.ParentSelectors) - 1

		for i, selector := range n.ParentSelectors {
			if i == last {
				fmt.Fprintf(w, "%s {\n", selector)
			} else {
				fmt.Fprintf(w, "%s,\n", selector)
			}
		}

		for _, child := range declarationNodes {
			child.Render(w)
		}

		fmt.Fprintf(w, "}\n")
	}

	for _, child := range n.Children {
		if child.GetType() != "declaration" {
			child.Render(w)
		}
	}

	fmt.Fprintln(w, "}")
}

func (n *atRuleNode) Dump(indent string) {
	fmt.Printf("%sAtRuleNode: %s\n", indent, n.Text)
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

type importNode struct {
	baseNode
	Text string
}

func (n *importNode) Render(w io.Writer) {
	fmt.Fprintf(w, "%s\n", n.Text)
}

func (n *importNode) Dump(indent string) {
	fmt.Printf("%sImportNode: %s\n", indent, n.Text)
}

type context struct {
	ParentContext *context
	Node          node
}

func (c *context) AddChild(node node) {
	c.Node.AddChild(node)
}

func newRootNode() *rootNode {
	n := rootNode{}
	n.Children = make([]node, 0)
	return &n
}

func newVariableNode(text string, lineNumber int) *variableNode {
	node := variableNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func newSelectorNode(selectors []string, lineNumber int) *selectorNode {
	n := selectorNode{}
	n.Children = make([]node, 0)
	n.Selectors = selectors
	n.LineNumber = lineNumber
	return &n
}

func newDeclarationNode(text string, lineNumber int) *declarationNode {
	node := declarationNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func newMixinNode(text string, lineNumber int) *mixinNode {
	node := mixinNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func newAtRuleNode(text string, lineNumber int) *atRuleNode {
	n := atRuleNode{Text: text}
	n.Children = make([]node, 0)
	n.LineNumber = lineNumber
	return &n
}

func newImportNode(text string, lineNumber int) *importNode {
	n := importNode{Text: text}
	n.LineNumber = lineNumber
	return &n
}

func newContext(node node, parentContext *context) *context {
	context := context{parentContext, node}
	return &context
}

func (p *parser) BuildTree(elements *elements) (*rootNode, error) {
	root := newRootNode()

	context := newContext(root, nil)

	for index, element := range elements.Items {
		elementType := element.ElementType
		text := element.Text
		lineNumber := element.LineNumber

		if elementType == typeVariable {
			variableNode := newVariableNode(text, lineNumber)
			context.AddChild(variableNode)
		}

		if elementType == typeDeclaration {
			declarations := splitDeclarations(text)

			for _, d := range declarations {
				declarationNode := newDeclarationNode(d, lineNumber)
				context.AddChild(declarationNode)
			}
		}

		if elementType == typeMixin {
			declarations := splitDeclarations(text)

			for _, d := range declarations {
				mixinNode := newMixinNode(d, lineNumber)
				context.AddChild(mixinNode)
			}
		}

		if elementType == typeImport {
			importNode := newImportNode(text, lineNumber)
			context.AddChild(importNode)
		}

		if elementType == typeOpenBrace {
			previousElement := elements.Items[index-1]
			previousType := previousElement.ElementType

			if previousType == typeSelector {
				selectors := getSelectors(elements, index-1)
				selectorNode := newSelectorNode(selectors, lineNumber)
				context.AddChild(selectorNode)
				context = newContext(selectorNode, context)
			} else if previousType == typeAtRule {
				atRuleNode := newAtRuleNode(previousElement.Text, lineNumber)
				context.AddChild(atRuleNode)
				context = newContext(atRuleNode, context)
			}

		}

		if elementType == typeCloseBrace {
			context = context.ParentContext
		}
	}

	return root, nil
}

func getSelectors(elements *elements, index int) []string {
	selectors := make([]string, 0)

	items := elements.Items
	element := &items[index]
	elementType := element.ElementType

	selectors = appendSelectors(selectors, element.Text)

	for index > 0 {
		index--
		element = &items[index]
		if element.ElementType != elementType {
			break
		}

		selectors = appendSelectors(selectors, element.Text)
	}

	slices.Reverse(selectors)

	return selectors
}

func appendSelectors(selectors []string, str string) []string {
	parts := strings.Split(str, ",")
	for _, part := range parts {
		selectors = append(selectors, strings.TrimSpace(part))
	}

	return selectors
}

func splitDeclarations(str string) []string {
	results := make([]string, 0)

	for {
		pos := strings.Index(str, ";")
		if pos < 0 {
			return results
		}

		s := strings.TrimSpace(str[0 : pos+1])
		str = str[pos+1:]

		results = append(results, s)
	}
}

func mergeSelectors(parentSelectors []string, childSelectors []string) []string {
	selectors := make([]string, 0)

	for _, childSelector := range childSelectors {
		if strings.Index(childSelector, "()") > 0 {
			continue
		}

		for _, parentSelector := range parentSelectors {
			var selector string
			if strings.Contains(childSelector, "&") {
				selector = strings.ReplaceAll(childSelector, "&", parentSelector)
			} else if parentSelector != "" {
				selector = parentSelector + " " + childSelector
			} else {
				selector = childSelector
			}

			selectors = append(selectors, selector)
		}
	}

	return selectors
}
