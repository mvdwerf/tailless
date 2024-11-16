package tailless

import "fmt"

func ResolveVariables(tree *RootNode) error {
	return RecursiveResolveVariables(tree, nil)
}

func RecursiveResolveVariables(node node, parentVariables *VariablesCollection) error {
	variables := NewVariablesCollection(parentVariables)
	variables.Read(node)

	err := node.ReplaceVariables(variables)
	if err != nil {
		return err
	}

	for _, child := range node.GetChildren() {
		err := RecursiveResolveVariables(child, variables)
		if err != nil {
			return err
		}
	}

	return nil
}

type VariablesCollection struct {
	Parent *VariablesCollection
	Items  map[string]string
}

func NewVariablesCollection(parent *VariablesCollection) *VariablesCollection {
	variables := VariablesCollection{Parent: parent}
	variables.Items = make(map[string]string)
	return &variables
}

func (v *VariablesCollection) Get(name string) string {
	value := v.Items[name]
	if value != "" {
		return value
	}

	if v.Parent != nil {
		return v.Parent.Get(name)
	}

	return ""
}

func (v *VariablesCollection) Read(n node) {
	for _, child := range n.GetChildren() {
		child.GetVariable(v)
	}
}

func (v *VariablesCollection) Set(name string, value string) {
	v.Items[name] = value
}

func (v *VariablesCollection) Dump() {
	for name, value := range v.Items {
		fmt.Printf("%s = %s\n", name, value)
	}
}
