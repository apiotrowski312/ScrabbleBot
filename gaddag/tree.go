package gaddag

import (
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

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

func (n *Node) addWord(word string) {
	word = strings.ToLower(word)
	for idx := range word {

		prefix := str_manipulator.Reverse(word[:len(word)-idx])
		sufix := word[len(word)-idx:]

		currentWord := prefix + "." + sufix
		currentNode := n
		for innerIndex, character := range currentWord {
			child, isOk := currentNode.get(character)

			if !isOk {
				isEndOfWord := innerIndex == len(currentWord)-1
				child = currentNode.add(character, Node{
					IsWord: isEndOfWord,
				})

				if isEndOfWord {
					break
				}
			}
			currentNode = child

		}
	}
}
