package gaddag

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

func (n *Node) addWord(word string) {
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
					isWord: isEndOfWord,
				})

				if isEndOfWord {
					break
				}
			}
			currentNode = child

		}
	}
}

// IsWordValid - check if provided string is marked as word in gaddag tree
func (n *Node) IsWordValid(word string) (bool, error) {
	currentNode := n
	var isOk bool
	for _, letter := range word {
		currentNode, isOk = currentNode.get(letter)

		if !isOk {
			return false, n.wordIsNotInDictionary(word)
		}
	}

	if currentNode.isWord {
		return true, nil
	}

	return false, n.wordIsNotInDictionary(word)
}

func (n *Node) wordIsNotInDictionary(word string) error {
	i := strings.Index(word, ".")

	if i < 0 {
		return fmt.Errorf("Word %v was passed in wrong format (no necessarry dot)", word)

	}
	processedWord := str_manipulator.Reverse(word[:i]) + word[i+1:]

	return fmt.Errorf("Word %v is not in dictionary", processedWord)
}

// CreateGraph - create gaddag tree structure based on file with all words, each starting from newline
func CreateGraph(filename string) (*Node, error) {
	root := &Node{
		children: map[rune]Node{},
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Fatal error while opening file: %v. Stacktrace: %v", filename, err)
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		word := sc.Text()
		root.addWord(word)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
		return nil, err
	}

	return root, nil
}
