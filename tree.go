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
	GetVariable(*VariablesCollection)
	ReplaceVariables(*VariablesCollection) error
	GetMixin(mixins)
	GetMixinName() string
	GetCopy() node
	GetLineNumber() int
	IsParentOf(node) bool
}

type BaseNode struct {
	Children   []node
	LineNumber int
	Hidden     bool
}

func (n *BaseNode) ExpandSelectors(parentSelectors []string) {

}

func (n *BaseNode) HideIfEmpty() bool {
	return false
}

func (n *BaseNode) GetChildren() []node {
	return n.Children
}

func (n *BaseNode) SetChildren(children []node) {
	n.Children = children
}

func (n *BaseNode) AddChild(child node) {
	n.Children = append(n.Children, child)
}

func (n *BaseNode) GetVariable(variables *VariablesCollection) {

}

func (n *BaseNode) ReplaceVariables(variables *VariablesCollection) error {
	return nil
}

func (n *BaseNode) GetMixin(mixins) {

}

func (n *BaseNode) GetMixinName() string {
	return ""
}

func (n *BaseNode) GetCopy() node {
	return nil
}

func (n *BaseNode) GetLineNumber() int {
	return n.LineNumber
}

func (n *BaseNode) Dump(indent string) {
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

func (n *BaseNode) Render(io.Writer) {

}

func (n *BaseNode) GetType() string {
	return ""
}

func (n *BaseNode) IsParentOf(node node) bool {
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

type RootNode struct {
	BaseNode
}

func ExpandSelectors(tree *RootNode) {
	for _, child := range tree.Children {
		child.ExpandSelectors([]string{""})
	}
}

func (n *RootNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *RootNode) Render(w io.Writer) {
	for _, child := range n.Children {
		child.Render(w)
	}
}

type VariableNode struct {
	BaseNode
	Text string
}

func (n *VariableNode) GetType() string {
	return "variable"
}

func (n *VariableNode) HideIfEmpty() bool {
	n.Hidden = true
	return true
}

func (n *VariableNode) GetVariable(variables *VariablesCollection) {
	parts := strings.Split(n.Text, ":")
	if len(parts) != 2 {
		return
	}

	name := strings.TrimPrefix(strings.TrimSpace(parts[0]), "@")
	value := strings.TrimSuffix(strings.TrimSpace(parts[1]), ";")

	variables.Set(name, value)
}

func (n *VariableNode) GetCopy() node {
	return NewVariableNode(n.Text, n.LineNumber)
}

func (n *VariableNode) Dump(indent string) {
	fmt.Printf("%sVariableNode: %s\n", indent, n.Text)
}

type SelectorNode struct {
	BaseNode
	Selectors       []string
	MergedSelectors []string
}

func (n *SelectorNode) ExpandSelectors(parentSelectors []string) {
	selectors := MergeSelectors(parentSelectors, n.Selectors)
	n.MergedSelectors = selectors
	for _, child := range n.Children {
		child.ExpandSelectors(selectors)
	}
}

func (n *SelectorNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *SelectorNode) GetMixin(mixins mixins) {
	selectors := n.Selectors
	if len(selectors) != 1 {
		return
	}

	name := strings.TrimSuffix(selectors[0], "()")

	mixins.Set(name, n)
}

func (n *SelectorNode) GetCopy() node {
	copy := NewSelectorNode(n.Selectors, n.LineNumber)

	for _, child := range n.Children {
		copy.Children = append(copy.Children, child.GetCopy())
	}

	return copy
}

func (n *SelectorNode) Dump(indent string) {
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

func (n *SelectorNode) Render(w io.Writer) {
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

type DeclarationNode struct {
	BaseNode
	Text string
}

func (n *DeclarationNode) GetType() string {
	return "declaration"
}

func (n *DeclarationNode) GetCopy() node {
	return NewDeclarationNode(n.Text, n.LineNumber)
}

func (n *DeclarationNode) Render(w io.Writer) {
	fmt.Fprintf(w, "  %s\n", n.Text)
}

func (n *DeclarationNode) ReplaceVariables(variables *VariablesCollection) error {
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

func (n *DeclarationNode) Dump(indent string) {
	fmt.Printf("%sDeclarationNode: %s\n", indent, n.Text)
}

type MixinNode struct {
	BaseNode
	Text string
}

func (n *MixinNode) ExpandSelectors(parentSelectors []string) {
	for _, child := range n.Children {
		child.ExpandSelectors(parentSelectors)
	}
}

func (n *MixinNode) GetMixinName() string {
	return strings.TrimSuffix(strings.TrimSuffix(n.Text, ";"), "()")
}

func (n *MixinNode) Render(w io.Writer) {
	for _, child := range n.Children {
		child.Render(w)
	}
}

func (n *MixinNode) Dump(indent string) {
	fmt.Printf("%sMixinNode: %s\n", indent, n.Text)
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

type AtRuleNode struct {
	BaseNode
	Text            string
	ParentSelectors []string
}

func (n *AtRuleNode) ExpandSelectors(parentSelectors []string) {
	n.ParentSelectors = parentSelectors
	for _, child := range n.Children {
		child.ExpandSelectors(parentSelectors)
	}
}

func (n *AtRuleNode) ReplaceVariables(variables *VariablesCollection) error {
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

func (n *AtRuleNode) HideIfEmpty() bool {
	isEmpty := true
	for _, child := range n.Children {
		if !child.HideIfEmpty() {
			isEmpty = false
		}
	}

	n.Hidden = isEmpty
	return isEmpty
}

func (n *AtRuleNode) Render(w io.Writer) {
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

func (n *AtRuleNode) Dump(indent string) {
	fmt.Printf("%sAtRuleNode: %s\n", indent, n.Text)
	for _, child := range n.Children {
		child.Dump(indent + "  ")
	}
}

type ImportNode struct {
	BaseNode
	Text string
}

func (n *ImportNode) Render(w io.Writer) {
	fmt.Fprintf(w, "%s\n", n.Text)
}

func (n *ImportNode) Dump(indent string) {
	fmt.Printf("%sImportNode: %s\n", indent, n.Text)
}

type Context struct {
	ParentContext *Context
	Node          node
}

func (c *Context) AddChild(node node) {
	c.Node.AddChild(node)
}

func NewRootNode() *RootNode {
	n := RootNode{}
	n.Children = make([]node, 0)
	return &n
}

func NewVariableNode(text string, lineNumber int) *VariableNode {
	node := VariableNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func NewSelectorNode(selectors []string, lineNumber int) *SelectorNode {
	n := SelectorNode{}
	n.Children = make([]node, 0)
	n.Selectors = selectors
	n.LineNumber = lineNumber
	return &n
}

func NewDeclarationNode(text string, lineNumber int) *DeclarationNode {
	node := DeclarationNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func NewMixinNode(text string, lineNumber int) *MixinNode {
	node := MixinNode{Text: text}
	node.LineNumber = lineNumber
	return &node
}

func NewAtRuleNode(text string, lineNumber int) *AtRuleNode {
	n := AtRuleNode{Text: text}
	n.Children = make([]node, 0)
	n.LineNumber = lineNumber
	return &n
}

func NewImportNode(text string, lineNumber int) *ImportNode {
	n := ImportNode{Text: text}
	n.LineNumber = lineNumber
	return &n
}

func NewContext(node node, parentContext *Context) *Context {
	context := Context{parentContext, node}
	return &context
}

func (p *parser) BuildTree(elements *elements) (*RootNode, error) {
	root := NewRootNode()

	context := NewContext(root, nil)

	for index, element := range elements.Items {
		elementType := element.ElementType
		text := element.Text
		lineNumber := element.LineNumber

		if elementType == Variable {
			variableNode := NewVariableNode(text, lineNumber)
			context.AddChild(variableNode)
		}

		if elementType == Declaration {
			declarations := splitDeclarations(text)

			for _, d := range declarations {
				declarationNode := NewDeclarationNode(d, lineNumber)
				context.AddChild(declarationNode)
			}
		}

		if elementType == Mixin {
			declarations := splitDeclarations(text)

			for _, d := range declarations {
				mixinNode := NewMixinNode(d, lineNumber)
				context.AddChild(mixinNode)
			}
		}

		if elementType == Import {
			importNode := NewImportNode(text, lineNumber)
			context.AddChild(importNode)
		}

		if elementType == OpenBrace {
			previousElement := elements.Items[index-1]
			previousType := previousElement.ElementType

			if previousType == Selector {
				selectors := getSelectors(elements, index-1)
				selectorNode := NewSelectorNode(selectors, lineNumber)
				context.AddChild(selectorNode)
				context = NewContext(selectorNode, context)
			} else if previousType == AtRule {
				atRuleNode := NewAtRuleNode(previousElement.Text, lineNumber)
				context.AddChild(atRuleNode)
				context = NewContext(atRuleNode, context)
			}

		}

		if elementType == CloseBrace {
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

func MergeSelectors(parentSelectors []string, childSelectors []string) []string {
	selectors := make([]string, 0)

	for _, childSelector := range childSelectors {
		if strings.Index(childSelector, "()") > 0 {
			continue
		}

		for _, parentSelector := range parentSelectors {
			var selector string
			if strings.Index(childSelector, "&") >= 0 {
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
