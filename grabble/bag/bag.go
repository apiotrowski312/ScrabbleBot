package bag

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

	if number > len(*b) {
		number = len(*b)
	}
	for i := 0; i < number; i++ {
		letterIndex := rand.Intn(len(*b))
		letters = append(letters, (*b)[letterIndex])
		(*b) = append((*b)[:letterIndex], (*b)[letterIndex+1:]...)
	}
	return letters
}

func (b *Bag) ChangeLetters(letters []rune) []rune {
	(*b) = append((*b), letters...)
	return b.DrawLetters(len(letters))
}

func (l LettersPoint) GetPoints(words, bonuses []string) int {
	allPoints := 0
	for oi, word := range words {
		wordBonus := 1
		points := 0
		for i, letter := range word {
			switch b := bonuses[oi][i]; b {
			case 'l':
				points += 2 * l[letter]
			case 'L':
				points += 3 * l[letter]
			case 'w':
				points += l[letter]
				wordBonus *= 2
			case 'W':
				points += l[letter]
				wordBonus *= 3
			case 's':
				points += l[letter]
				wordBonus *= 2
			default:
				points += l[letter]
			}

		}
		allPoints += points * wordBonus
	}

	return allPoints
}
