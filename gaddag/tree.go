package gaddag

// Node - struct on which whole gaddag is build.
type Node struct {
	children map[rune]Node
	isWord   bool
	letter   rune
}

// get return new Node and boolen isOk
func (n *Node) get(path rune) (*Node, bool) {
	child, ok := n.children[path]
	return &child, ok
}

func (n Node) add(letter rune, child Node) *Node {
	if child.children == nil {
		child.children = map[rune]Node{}
		child.letter = letter
	}

	n.children[letter] = child
	return &child
}
