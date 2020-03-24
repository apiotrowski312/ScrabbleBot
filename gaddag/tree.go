package gaddag

type node struct {
	children map[rune]node
	isWord   bool
}

func (n *node) get(path rune) (*node, bool) {
	child, ok := n.children[path]
	return &child, ok
}

func (n node) add(path rune, child node) *node {
	if child.children == nil {
		child.children = map[rune]node{}
	}

	n.children[path] = child
	return &child
}
