package player

import "fmt"

type Player struct {
	Name   string
	Points int
	Rack   []rune
}

func CreatePlayer(name string) Player {
	return Player{Name: name}
}

func (p *Player) UpdateRack(lettersToRemove, lettersToAdd []rune) error {
	if err := p.removeFromRack(lettersToRemove); err != nil {
		return err
	}

	p.addToRack(lettersToAdd)

	return nil
}

func (p *Player) removeFromRack(lettersToRemove []rune) error {
	for _, l := range lettersToRemove {
		letterInRack := false
		for i := range p.Rack {
			if p.Rack[i] == l {
				p.Rack = append(p.Rack[:i], p.Rack[i+1:]...)
				letterInRack = true
				break
			}
		}
		if letterInRack == false {
			return fmt.Errorf("Letter %s is not in your rack. It cannot be removed", string(l))
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
