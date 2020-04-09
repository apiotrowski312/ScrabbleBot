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

func (b board) collectOtherWordsAndTilesHorizontal(word string, startCord [2]int) ([]string, []string) {
	words := []string{}
	tileTypes := []string{}

	for index, letter := range word {
		if b[startCord[0]][startCord[1]+index].Letter == letter {
			continue
		}
		currentWord := string(letter)
		currentTile := string(b[startCord[0]][startCord[1]+index].TileType)

		// Up side
		innerIndex := startCord[0] - 1
		for innerIndex >= 0 {
			currentLetter := b[innerIndex][startCord[1]+index].Letter
			if currentLetter == rune(0) {
				break
			}
			currentWord += string(currentLetter)
			innerIndex--
		}

		currentWord += "."
		// Down side
		innerIndex = startCord[0] + 1
		for innerIndex <= len(b) {
			currentLetter := b[innerIndex][startCord[1]+index].Letter
			if currentLetter == rune(0) {
				break
			}
			currentWord += string(currentLetter)
			innerIndex++
		}

		if len(currentWord) > 2 {
			words = append(words, currentWord)
			tileTypes = append(tileTypes, currentTile)
		}
	}
	return words, tileTypes
}

func (b board) collectOtherWordsAndTilesVertical(word string, startCord [2]int) ([]string, []string) {
	words := []string{}
	tileTypes := []string{}

	for index, letter := range word {
		if b[startCord[0]+index][startCord[1]].Letter == letter {
			continue
		}
		currentWord := string(letter)
		currentTile := string(b[startCord[0]+index][startCord[1]].TileType)

		// Left side
		innerIndex := startCord[1] - 1
		for innerIndex >= 0 {
			currentLetter := b[startCord[0]+index][innerIndex].Letter
			if currentLetter == rune(0) {
				break
			}
			currentWord += string(currentLetter)
			innerIndex--
		}

		currentWord += "."
		// Right side
		innerIndex = startCord[1] + 1
		for innerIndex <= len(b) {
			currentLetter := b[startCord[0]+index][innerIndex].Letter
			if currentLetter == rune(0) {
				break
			}
			currentWord += string(currentLetter)
			innerIndex++
		}

		if len(currentWord) > 2 {
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

func (b board) tileUnderLayedWord(word string, startCord [2]int, horizontal bool) string {
	currentTile := ""
	for index := range word {
		if horizontal {
			currentTile += string(b.getTileType([2]int{startCord[0], startCord[1] + index}))
		} else {
			currentTile += string(b.getTileType([2]int{startCord[0] + index, startCord[1]}))
		}
	}
	return currentTile
}

func (b board) getTileType(cord [2]int) rune {
	if b[cord[0]][cord[1]].Letter != rune(0) {
		return '0'
	}
	return b[cord[0]][cord[1]].TileType
}

func (b board) collectAllUsedWords(word string, startCord [2]int, horizontal bool) ([]string, []string) {
	words := []string{}
	tiles := []string{}
	if horizontal {
		words, tiles = b.collectOtherWordsAndTilesHorizontal(word, startCord)
	} else {
		words, tiles = b.collectOtherWordsAndTilesVertical(word, startCord)
	}

	mainWordTiles := b.tileUnderLayedWord(word, startCord, horizontal)

	words = append(words, word)
	tiles = append(tiles, mainWordTiles)

	return words, tiles

}
