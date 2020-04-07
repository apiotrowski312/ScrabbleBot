package gaddag

// Node - struct on which whole gaddag is build.
type Node struct {
	children map[rune]Node
	isWord   bool
}

func (n *Node) get(path rune) (*Node, bool) {
	child, ok := n.children[path]
	return &child, ok
}

func (n Node) add(path rune, child Node) *Node {
	if child.children == nil {
		child.children = map[rune]Node{}
	}

	n.children[path] = child
	return &child
}
