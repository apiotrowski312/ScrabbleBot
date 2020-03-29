package gaddag

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

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

func (n *Node) IsWordValid(word string) (bool, error) {
	currentNode := n
	var isOk bool
	for _, letter := range word {
		currentNode, isOk = currentNode.get(letter)

		if !isOk {
			return false, errors.New(fmt.Sprintf("Word %v is not in dictionary", word))
		}
	}

	if currentNode.isWord {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Word %v is not in dictionary", word))
	}
}

func CreateGraph(filename string) (*Node, error) {
	root := &Node{
		children: map[rune]Node{},
	}

	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		word := sc.Text()
		root.addWord(word)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil, err
	}

	return root, nil
}
