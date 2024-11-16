package tailless

import "fmt"

type mixins interface {
	Get(string) *selectorNode
	Set(string, *selectorNode)
}

func resolveMixins(tree *rootNode) error {
	return recursiveResolveMixins(tree, nil, newTailwindCollection())
}

func recursiveResolveMixins(n node, parentMixins mixins, twMixins mixins) error {
	mixins := newMixinsCollection(parentMixins)
	mixins.Read(n)

	newChildren := make([]node, 0)

	for _, child := range n.GetChildren() {
		name := child.GetMixinName()
		if name == "" {
			newChildren = append(newChildren, child)

			err := recursiveResolveMixins(child, mixins, twMixins)
			if err != nil {
				return err
			}

			continue
		}

		mixin := twMixins.Get(name)
		if mixin == nil {
			mixin = mixins.Get(name)
		}

		if mixin == nil {
			return fmt.Errorf("Line %d: Mixin '%s' not found", child.GetLineNumber(), name)
		}

		if mixin.IsParentOf(child) {
			return fmt.Errorf("Line %d: Invalid parent mixin '%s'", child.GetLineNumber(), name)
		}

		newChildren = append(newChildren, mixin.GetCopy().GetChildren()...)
	}

	n.SetChildren(newChildren)

	return nil
}

type mixinsCollection struct {
	Parent mixins
	Items  map[string]*selectorNode
}

func newMixinsCollection(parent mixins) *mixinsCollection {
	mixins := mixinsCollection{Parent: parent}
	mixins.Items = make(map[string]*selectorNode)
	return &mixins
}

func (m *mixinsCollection) Get(name string) *selectorNode {
	node := m.Items[name]
	if node != nil {
		return node
	}

	if m.Parent != nil {
		return m.Parent.Get(name)
	}

	return nil
}

func (m *mixinsCollection) Read(n node) {
	for _, child := range n.GetChildren() {
		child.GetMixin(m)
	}
}

func (m *mixinsCollection) Set(name string, node *selectorNode) {
	m.Items[name] = node
}
