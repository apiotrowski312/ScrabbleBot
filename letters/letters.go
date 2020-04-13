package letters

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

type LetterValue map[rune]int
type TileBag []rune

func LoadTilesFromFile(filename string) (*TileBag, *LetterValue, error) {
	csvFile, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Fatal error when opening file %v with letters number and value. Stacktrace: %v", filename, err)
		return nil, nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.Comment = '#'

	// This is created to omit first line of csv (headers)
	if _, err = r.Read(); err != nil {
		log.Fatalf("Fatal error while reading file %v. Wrong headers. Stacktrace: %v", filename, err)
		return nil, nil, err
	}

	tB := TileBag{}
	lV := LetterValue{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		letter := []rune(strings.TrimSpace(strings.ToLower(record[0])))[0]

		lV[letter], err = strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			log.Fatalf("Fatal error while reading file %v. Wrong line content: %v. Stacktrace: %v", filename, record, err)
			return nil, nil, err
		}

		numberOfTiles, err := strconv.Atoi(strings.TrimSpace(record[2]))

		for i := 0; i < numberOfTiles; i++ {
			tB = append(tB, letter)
		}

		if err != nil {
			log.Fatalf("Fatal error while reading file %v. Wrong line content: %v. Stacktrace: %v", filename, record, err)
			return nil, nil, err
		}
	}
	return &tB, &lV, nil
}

// CountPoints - count points based on passed string slices
func (l LetterValue) CountPoints(words []string, tileTypes []string) int {
	points := 0
	for index, word := range words {
		wordPoints := 0
		multiplayer := 1
		word = str_manipulator.RemoveCharacters(word, ".")
		for innerIndex, letter := range word {
			currentTile := '0'
			if len(tileTypes[index])-1 >= innerIndex {
				currentTile = rune(tileTypes[index][innerIndex])
			}

			switch tile := currentTile; tile {
			case '0':
				wordPoints += l[letter]
			case 's':
				wordPoints += l[letter]
			case 'l':
				wordPoints += 2 * l[letter]
			case 'L':
				wordPoints += 3 * l[letter]
			case 'w':
				multiplayer *= 2
				wordPoints += l[letter]
			case 'W':
				multiplayer *= 3
				wordPoints += l[letter]
			}
		}
		points += multiplayer * wordPoints
	}
	return points
}

// DrawTiles returns rune slice with random draw letters.
// After each draw, letter is deleted from TileBag.
func (tb *TileBag) DrawTiles(numberOfTiles int) []rune {
	tiles := []rune{}

	for i := 0; i < numberOfTiles; i++ {
		if len((*tb)) <= 0 {
			return tiles
		}
		index := rand.Intn(len((*tb)))

		tiles = append(tiles, (*tb)[index])
		(*tb) = append((*tb)[:index], (*tb)[index+1:]...)
	}

	return tiles
}
