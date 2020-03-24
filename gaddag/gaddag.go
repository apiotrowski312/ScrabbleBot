package gaddag

import "github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"

func (n *node) addWord(word string) {
	for idx := range word {

		prefix := str_manipulator.Reverse(word[:len(word)-idx])
		sufix := word[len(word)-idx:]

		currentWord := prefix + "." + sufix
		currentNode := n
		for innerIndex, character := range currentWord {
			child, isOk := currentNode.get(character)

			if !isOk {
				isEndOfWord := innerIndex == len(currentWord)-1
				child = currentNode.add(character, node{
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
