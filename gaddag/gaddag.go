package gaddag

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func (n Node) FindAllWords(hook rune, letters []rune) []string {
	newLetters := append(letters, '.')
	words := n.getAllOk(hook, newLetters)
	return removeDuplicatesUnordered(words)
}

func (n Node) getAllOk(currentLetter rune, lettersToGo []rune) []string {
	hookNode, isOk := n.get(currentLetter)

	if !isOk {
		return nil
	}

	partialWords := []string{}
	if hookNode.IsWord {
		partialWords = append(partialWords, string(currentLetter))
	}

	for i, l := range lettersToGo {
		lettersForIteration := append(append([]rune{}, lettersToGo[:i]...), lettersToGo[i+1:]...)

		newWords := hookNode.getAllOk(l, lettersForIteration)
		for _, w := range newWords {
			partialWords = append(partialWords, string(currentLetter)+w)
		}
	}

	// fmt.Println(string(currentLetter), ":", partialWords)

	return partialWords
}

// IsWordValid - check if provided string is marked as word in gaddag tree
func (n *Node) IsWordValid(word string) (bool, error) {
	i := strings.Index(word, ".")
	if i < 0 {
		word = word[:1] + "." + word[1:]
	}

	currentNode := n
	var isOk bool
	for _, letter := range word {
		currentNode, isOk = currentNode.get(letter)

		if !isOk {
			return false, n.wordIsNotInDictionary(word)
		}
	}

	if currentNode.IsWord {
		return true, nil
	}

	return false, n.wordIsNotInDictionary(word)
}

func (n *Node) wordIsNotInDictionary(word string) error {
	i := strings.Index(word, ".")
	processedWord := str_manipulator.Reverse(word[:i]) + word[i+1:]

	return fmt.Errorf("Word %v is not in dictionary", processedWord)
}

// CreateGraph - create gaddag tree structure based on file with all words, each starting from newline
func CreateGraph(filename string) (*Node, error) {
	root := &Node{
		Children: map[rune]Node{},
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
