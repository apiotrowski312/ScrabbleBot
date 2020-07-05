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
	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// FindAllWords - pass hookIndex, row and letters.
// Function will find all words based od available letters and already existring one.
func (n Node) FindAllWords(hookIndex int, row []rune, letters []rune) []string {
	newLetters := append(letters, '.')

	words := n.getAllOk(row[hookIndex], hookIndex, newLetters, row, hookIndex)
	return removeDuplicatesUnordered(words)
}

// TODO: If there is a letter on right, do not append it, it wont be valid word anyway
func (n Node) getAllOk(currentLetter rune, letterIndex int, lettersToUse []rune, row []rune, hookIndex int) []string {

	if letterIndex == -1 && currentLetter != '.' {
		return nil
	}

	hookNode, isOk := n.get(currentLetter)
	if !isOk {
		return nil
	}

	partialWords := []string{}
	if hookNode.IsWord {
		partialWords = append(partialWords, string(currentLetter))
	}

	var newLetterIndex int
	if currentLetter == '.' {
		newLetterIndex = hookIndex + 1
	} else if letterIndex <= hookIndex {
		newLetterIndex = letterIndex - 1
	} else {
		newLetterIndex = letterIndex + 1
	}

	if newLetterIndex == len(row) {
		if hookNode.IsWord {
			return partialWords
		}
		return nil
	}

	// FIXME: There is no point in checking for smth else than dot in -1
	if newLetterIndex >= 0 && row[newLetterIndex] != rune(0) {
		newWords := hookNode.getAllOk(row[newLetterIndex], newLetterIndex, lettersToUse, row, hookIndex)
		for _, w := range newWords {
			partialWords = append(partialWords, string(currentLetter)+w)
		}
	} else {
		for i, l := range lettersToUse {
			lettersForIteration := append(append([]rune{}, lettersToUse[:i]...), lettersToUse[i+1:]...)
			newWords := hookNode.getAllOk(l, newLetterIndex, lettersForIteration, row, hookIndex)

			for _, w := range newWords {
				partialWords = append(partialWords, string(currentLetter)+w)
			}
		}
	}

	return partialWords
}

// IsWordValid - check if provided string is marked as a word in gaddag tree
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
