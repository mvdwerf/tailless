package tailless

import "fmt"

func resolveVariables(tree *rootNode) error {
	return recursiveResolveVariables(tree, nil)
}

func recursiveResolveVariables(node node, parentVariables *variablesCollection) error {
	variables := newVariablesCollection(parentVariables)
	variables.Read(node)

	err := node.ReplaceVariables(variables)
	if err != nil {
		return err
	}

	for _, child := range node.GetChildren() {
		err := recursiveResolveVariables(child, variables)
		if err != nil {
			return err
		}
	}

	return nil
}

type variablesCollection struct {
	Parent *variablesCollection
	Items  map[string]string
}

func newVariablesCollection(parent *variablesCollection) *variablesCollection {
	variables := variablesCollection{Parent: parent}
	variables.Items = make(map[string]string)
	return &variables
}

func (v *variablesCollection) Get(name string) string {
	value := v.Items[name]
	if value != "" {
		return value
	}

	if v.Parent != nil {
		return v.Parent.Get(name)
	}

	return ""
}

func (v *variablesCollection) Read(n node) {
	for _, child := range n.GetChildren() {
		child.GetVariable(v)
	}
}

func (v *variablesCollection) Set(name string, value string) {
	v.Items[name] = value
}

func (v *variablesCollection) Dump() {
	for name, value := range v.Items {
		fmt.Printf("%s = %s\n", name, value)
	}
}
