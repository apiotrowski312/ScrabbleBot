package gaddag

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

const blankLetters = "EAIONRTLSUUDGBCMPFHVWYKJXQZ"

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
// Function will find all words based od available letters and already existing one.
// Passed row have to have at least one word inside. This algorithm works only with existing hook.
func (n Node) FindAllWords(hookIndex int, row []rune, letters []rune) []string {
	newLetters := append([]rune{}, append(letters, '.')...)

	sort.Slice(newLetters, func(i, j int) bool {
		return newLetters[i] < newLetters[j]
	})

	var words []string
	if i := strings.Index(string(newLetters), "_"); i == -1 {
		words = n.getAllOk(row[hookIndex], hookIndex, newLetters, row, hookIndex)
	} else {
		lettersExceptBlanks := newLetters[:i]
		numOfBlanks := len(newLetters) - i

		permutation := []rune{}
		for a := 1; a <= numOfBlanks; a++ {
			permutation = append(permutation, 'a')
		}

		var err error
		for {
			permutationLetters := append(append([]rune{}, lettersExceptBlanks...), permutation...)
			wordsFromPermutation := n.getAllOk(row[hookIndex], hookIndex, permutationLetters, row, hookIndex)
			words = append(words, wordsFromPermutation...)
			if permutation, err = GetNextPermutation(permutation); err != nil {
				break
			}
		}
	}

	return removeDuplicatesUnordered(words)
}

// getAllOk return string array with all possible combinations of words with letters
func (n Node) getAllOk(currentLetter rune, letterIndex int, lettersToUse []rune, row []rune, hookIndex int) []string {

	if letterIndex == -1 && currentLetter != '.' {
		return nil
	}

	hookNode, isOk := n.get(unicode.ToUpper(currentLetter))
	if !isOk {
		return nil
	}

	partialWords := []string{}
	if hookNode.IsWord {
		if letterIndex == len(row)-1 || hookIndex == len(row)-1 {
			partialWords = append(partialWords, string(currentLetter))
		} else if letterIndex < hookIndex && row[hookIndex+1] == rune(0) {
			partialWords = append(partialWords, string(currentLetter))
		} else if letterIndex >= hookIndex && row[letterIndex+1] == rune(0) {
			partialWords = append(partialWords, string(currentLetter))
		}
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

	if newLetterIndex >= 0 && row[newLetterIndex] != rune(0) {
		newWords := hookNode.getAllOk(row[newLetterIndex], newLetterIndex, lettersToUse, row, hookIndex)
		for _, w := range newWords {
			partialWords = append(partialWords, string(currentLetter)+w)
		}
	} else if newLetterIndex == -1 && lettersToUse[0] == '.' {
		lettersForIteration := append(append([]rune{}, lettersToUse[:0]...), lettersToUse[1:]...)
		newWords := hookNode.getAllOk('.', newLetterIndex, lettersForIteration, row, hookIndex)

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
	if i == -1 {
		word = word[:1] + "." + word[1:]
	}

	word = strings.ToUpper(word)

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
	return fmt.Errorf("Word %v is not in dictionary", NormalizeWord(word))
}

// NormalizeWord returns word converted from gaddag notation to normal
func NormalizeWord(word string) string {
	i := strings.Index(word, ".")
	return str_manipulator.Reverse(word[:i]) + word[i+1:]
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

func GetNextPermutation(permutation []rune) ([]rune, error) {
	for i := range permutation {
		if permutation[i] == 'z' {
			for _, p := range permutation[i+1:] {
				if p != 'z' {
					permutation[i] = p + 1
					break
				}
			}
		} else {
			permutation[i]++
			return permutation, nil
		}
	}

	return []rune{}, fmt.Errorf("Next permutation not exists")
}
