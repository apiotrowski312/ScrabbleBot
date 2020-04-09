package game

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"
)

type letterValue map[rune]int
type tileBag map[rune]int

func loadTilesFromFile(filename string) (*tileBag, *letterValue, error) {
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

	tB := tileBag{}
	lV := letterValue{}

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

		tB[letter], err = strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			log.Fatalf("Fatal error while reading file %v. Wrong line content: %v. Stacktrace: %v", filename, record, err)
			return nil, nil, err
		}
	}
	return &tB, &lV, nil
}

func (l letterValue) countPoints(words []string, tileTypes []string) int { // TODO: remove ifology
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
