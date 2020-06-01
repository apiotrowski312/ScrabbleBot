package grabble

import (
	"math/rand"
	"time"
)

type Bag []rune
type LettersPoint map[rune]int

func CreateBag(tiles []rune) Bag {
	return Bag(tiles)
}

func CreateLettersPoint(lp map[rune]int) LettersPoint {
	return LettersPoint(lp)
}

func (b *Bag) DrawLetters(number int) []rune {
	var letters []rune
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		letterIndex := rand.Intn(len((*b)))
		letters = append(letters, (*b)[letterIndex])
		(*b) = append((*b)[:letterIndex], (*b)[letterIndex+1:]...)
	}
	return letters
}

// TODO: Put letters back to the bag when user change letters
