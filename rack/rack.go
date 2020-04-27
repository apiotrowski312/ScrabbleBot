package rack

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Rack struct {
	letters  []rune
	rackSize int
}

type ExportedRack struct {
	Letters  []rune
	RackSize int
}

func CreateRack(rS int) Rack {
	return Rack{rackSize: rS}
}

func (r Rack) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(ExportedRack{
		Letters:  r.letters,
		RackSize: r.rackSize,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (r *Rack) UnmarshalJSON(jsonBytes []byte) error {
	var exportedRack ExportedRack
	if err := json.Unmarshal(jsonBytes, &exportedRack); err != nil {
		return err
	}

	(*r).letters = exportedRack.Letters
	(*r).rackSize = exportedRack.RackSize
	return nil
}

func (r *Rack) RemoveFromRack(toRemove []rune) error {

	if len((*r).letters) == 0 {
		return errors.New("Empty rack, you cannot remove any letter from it")
	}

	for _, toR := range toRemove {
		for i, ra := range (*r).letters {
			if ra == toR {
				(*r).letters = append((*r).letters[:i], (*r).letters[i+1:]...)
				break
			}

			if i == len((*r).letters)-1 {
				return fmt.Errorf("There is no tile %v in rack", string(toR))
			}
		}
	}
	return nil
}

func (r *Rack) AddToRack(toAdd []rune) error {
	if len((*r).letters)+len(toAdd) > r.rackSize {
		return errors.New("Rack will be overfilled")
	}

	(*r).letters = append((*r).letters, toAdd...)

	return nil
}

func (r *Rack) AreThereLetters(toCheck []rune) (bool, error) {
	index := 0
	for _, toC := range toCheck {
		for _, ra := range (*r).letters {
			if toC == ra {
				index++
			}
		}
	}

	if index >= len(toCheck) {
		return true, nil
	}

	return false, errors.New("There is no such letters")
}
