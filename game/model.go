package game

import (
	"encoding/json"

	"github.com/apiotrowski312/scrabbleBot/board"
	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/letters"
	"github.com/apiotrowski312/scrabbleBot/rack"
)

type player struct {
	name   string
	rack   rack.Rack
	points int
}

type jsonPlayer struct {
	Name   string
	Rack   rack.Rack
	Points int
}

func (p player) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(jsonPlayer{
		Points: p.points,
		Name:   p.name,
		Rack:   p.rack,
	})

	if err != nil {
		return nil, err
	}
	return j, nil
}

func (p *player) UnmarshalJSON(jsonBytes []byte) error {
	var exportedPlayer jsonPlayer
	if err := json.Unmarshal(jsonBytes, &exportedPlayer); err != nil {
		return err
	}

	p.name = exportedPlayer.Name
	p.points = exportedPlayer.Points
	p.rack = exportedPlayer.Rack

	return nil
}

type game struct {
	bag          letters.TileBag
	letterValues letters.LetterValue
	board        board.Board
	dictionary   gaddag.Node
	players      []player
	round        int
}

type jsonGame struct {
	Bag          letters.TileBag
	LetterValues letters.LetterValue
	Board        board.Board
	Dictionary   gaddag.Node
	Players      []player
	Round        int
}

func (g game) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(jsonGame{
		Bag:          g.bag,
		LetterValues: g.letterValues,
		Board:        g.board,
		Dictionary:   g.dictionary,
		Players:      g.players,
		Round:        g.round,
	})

	if err != nil {
		return nil, err
	}
	return j, nil
}

func (g *game) UnmarshalJSON(jsonBytes []byte) error {
	var exportedGame jsonGame
	if err := json.Unmarshal(jsonBytes, &exportedGame); err != nil {
		return err
	}

	g.bag = exportedGame.Bag
	g.letterValues = exportedGame.LetterValues
	g.board = exportedGame.Board
	g.dictionary = exportedGame.Dictionary
	g.players = exportedGame.Players
	g.round = exportedGame.Round

	return nil
}
