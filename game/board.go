package game

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	tileType = `[0lLwWs]`
)

type board [][]tile

type tile struct {
	Letter   rune
	TileType rune
}

func loadBoardFromFile(filename string) (*board, error) {

	boardFile, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Fatal error when opening file with board template: %v", err)
		return nil, err
	}
	defer boardFile.Close()

	sc := bufio.NewScanner(boardFile)

	board := board{}

	fileLine := 0
	for sc.Scan() {
		fileLine++
		word := strings.TrimSpace(sc.Text())
		if word == "" || word[0] == '#' {
			continue
		}

		if matched, _ := regexp.MatchString(tileType, word); !matched {
			log.Fatalf("Fatal error while loading board to struct. Error in %v file. Board scheme with wrong character in scheme %v/ Line %v", filename, word, fileLine)
			return nil, errors.New("Wrong board scheme")
		}

		row := []tile{}

		for _, l := range word {
			row = append(row, tile{TileType: l})
		}
		board = append(board, row)

	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Fatal error with scanning file: %v", err)
		return nil, err
	}

	return &board, nil
}

func (b board) isWordInProperPlace(word string, startCord [2]int, horizontal bool) (bool, error) {

	isAddedCorectly := false

	for index, letter := range word {
		var currentTile tile
		if horizontal {
			currentTile = b[startCord[0]][startCord[1]+index]
		} else {
			currentTile = b[startCord[0]+index][startCord[1]]
		}
		if currentTile.Letter != rune(0) && currentTile.Letter != letter {
			return false, errors.New("You can't overwrite letter")
		}
		if currentTile.Letter == letter {
			isAddedCorectly = true
		} else if currentTile.TileType == 's' {
			return true, nil
		}
	}
	if isAddedCorectly {
		return true, nil
	}
	return false, errors.New("There is no hooks. Wrong place")
}

func (b board) collectAllUsedWords(word string, startCord [2]int, horizontal bool) ([]string, []string) {
	// TODO: refactor (remove ifology)
	words := []string{word}
	tileTypes := []string{}

	currentTile := ""
	for index := range word {
		if horizontal {
			currentTile += string(b[startCord[0]][startCord[1]+index].TileType)
		} else {
			currentTile += string(b[startCord[0]+index][startCord[1]].TileType)
		}
	}

	tileTypes = append(tileTypes, currentTile)

	for index, letter := range word {
		currentWord := string(letter)
		if horizontal && b[startCord[0]][startCord[1]+index].Letter == letter {
			continue
		} else if !horizontal && b[startCord[0]+index][startCord[1]].Letter == letter {
			continue
		}

		if horizontal {
			innerIndex := 1
			currentTile = string(b[startCord[0]][startCord[1]+index].TileType)
			for {
				if startCord[0]-innerIndex == -1 {
					break
				}
				currentLetter := b[startCord[0]-innerIndex][startCord[1]+index].Letter
				if currentLetter == rune(0) {
					break
				}
				currentWord += string(currentLetter)
				innerIndex++
			}
			currentWord += "."
			innerIndex = 1
			for {
				if startCord[0]+innerIndex == len(b) { // TODO: Should be handled better (what if board is not a cube)
					break
				}

				currentLetter := b[startCord[0]+innerIndex][startCord[1]+index].Letter
				if currentLetter == rune(0) {
					break
				}
				currentWord += string(currentLetter)
				innerIndex++
			}
		} else {
			innerIndex := 0
			currentTile = string(b[startCord[0]][startCord[1]+index].TileType)
			for {
				if startCord[0]-innerIndex == -1 {
					break
				}
				currentLetter := b[startCord[0]+index][startCord[1]-innerIndex].Letter
				if currentLetter != ' ' {
					break
				}
				currentWord += string(currentLetter)
				innerIndex++
			}
			innerIndex = 1
			for {
				if startCord[0]+innerIndex == len(b[startCord[0]+index]) {
					break
				}

				currentLetter := b[startCord[0]+index][startCord[1]+innerIndex].Letter
				if currentLetter == ' ' {
					break
				}
				currentWord += string(currentLetter)
				innerIndex++
			}
		}
		if currentWord != string(letter)+"." {
			words = append(words, currentWord)
			tileTypes = append(tileTypes, currentTile)
		}
	}

	return words, tileTypes
}

func (b board) placeWord(word string, startCord [2]int, horizontal bool) {
	for index, letter := range word {
		tile := &tile{}
		if horizontal {
			tile = &b[startCord[0]][startCord[1]+index]
		} else {
			tile = &b[startCord[0]+index][startCord[1]]
		}
		tile.Letter = letter
	}
}
