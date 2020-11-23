package player

import (
	"fmt"
	"unicode"
)

type Player struct {
	Name     string
	Points   int
	Rack     []rune
	Strategy string
}

func CreatePlayer(name string) Player {
	return Player{
		Name:   name,
		Points: 0,
	}
}

func (p *Player) UpdateRack(lettersToRemove, lettersToAdd []rune) error {
	if err := p.removeFromRack(lettersToRemove); err != nil {
		return err
	}

	p.addToRack(lettersToAdd)

	return nil
}

func (p *Player) removeFromRack(lettersToRemove []rune) error {
	if err := p.AreLettersInRack(lettersToRemove); err != nil {
		return err
	}

	for _, l := range lettersToRemove {
		if unicode.IsLower(l) {
			l = '_'
		}

		for i := range p.Rack {
			if p.Rack[i] == l {
				p.Rack = append(p.Rack[:i], p.Rack[i+1:]...)
				break
			}
		}
	}

	return nil
}

func (p *Player) addToRack(lettersToAdd []rune) {
	p.Rack = append(p.Rack, lettersToAdd...)
}

func (p *Player) AddPoints(points int) {
	p.Points += points
}

func (p *Player) MinusPoints(points int) {
	p.Points -= points
}

// AreLettersInRack - iterate over all letters and check if all are in user Rack
func (p *Player) AreLettersInRack(letters []rune) error {
	alreadyChecked := make(map[int]bool)

	for _, l := range letters {
		if unicode.IsLower(l) {
			l = '_'
		}

		foundLetter := false

		for i, r := range p.Rack {
			if isOk := alreadyChecked[i]; !isOk && l == r {
				alreadyChecked[i] = true
				foundLetter = true
				break
			}
		}

		if foundLetter == false {
			return fmt.Errorf("Letter %v(%v) is not in your rack", string(l), l)
		}
	}
	return nil
}
