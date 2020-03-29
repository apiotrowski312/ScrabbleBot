package game

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type letterValue map[rune]int
type tileBag map[rune]int

func loadTilesFromFile(filename string) (*tileBag, *letterValue, error) {
	csvFile, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil, nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.Comment = '#'

	// This is created to omit first line of csv (headers)
	if _, err = r.Read(); err != nil {
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
			return nil, nil, err
		}

		tB[letter], err = strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			return nil, nil, err
		}
	}
	return &tB, &lV, nil
}
