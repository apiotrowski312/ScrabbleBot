package gaddag

import (
	"encoding/gob"
	"os"
	"reflect"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

// Node - struct on which whole gaddag is build.
type Node struct {
	IsWord   bool
	Children map[rune]*Node
}

// get return child node for provided rune if exist.
func (n *Node) get(path rune) (*Node, bool) {
	child, ok := n.Children[path]
	return child, ok
}

func (n Node) add(letter rune, child Node) *Node {
	if child.Children == nil {
		child.Children = map[rune]*Node{}
	}

	n.Children[letter] = &child
	return &child
}

func (n *Node) addWord(word string) {
	word = strings.ToUpper(word)

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

func (n *Node) morphTreeToGADDAG() {
	n._morphTreeToGADDAG(n)
}

func (n *Node) _morphTreeToGADDAG(compareTo *Node) {
	for k, child := range n.Children {
		child._morphTreeToGADDAG(compareTo)

		if childrenPointer := n.findSameNode(child); childrenPointer != nil {
			n.Children[k] = childrenPointer
		}
	}
}

func (n *Node) findSameNode(lookFor *Node) *Node {
	for _, cn := range n.Children {
		if nPointer := cn.findSameNode(lookFor); nPointer != nil {
			return nPointer
		}
	}

	if n == lookFor {
		return lookFor
	} else if reflect.DeepEqual(n, lookFor) {
		return n
	}

	return nil
}

func (n Node) SaveToFile(filePath string) {
	dataFile, err := os.Create(filePath)

	if err != nil {
		panic("Couldn't open the file: " + err.Error())
	}
	defer dataFile.Close()

	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(n)
}

func LoadFromFile(filePath string) Node {
	var root Node
	dataFile, err := os.Open(filePath)

	if err != nil {
		panic("Couldn't open the file: " + err.Error())
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&root)

	if err != nil {
		panic("Couldn't decode the file: " + err.Error())
	}

	return root
}
