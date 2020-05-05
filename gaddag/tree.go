package gaddag

// Node - struct on which whole gaddag is build.
type Node struct {
	Letter   rune
	IsWord   bool
	Children map[rune]Node
}

// get return new Node and boolen isOk
func (n *Node) get(path rune) (*Node, bool) {
	child, ok := n.Children[path]
	return &child, ok
}

func (n Node) add(letter rune, child Node) *Node {
	if child.Children == nil {
		child.Children = map[rune]Node{}
		child.Letter = letter
	}

	n.Children[letter] = child
	return &child
}
