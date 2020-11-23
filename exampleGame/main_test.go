package main

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func TestGame(t *testing.T) {
	game := grabble.CreateDefaultGame([]string{"Bot 1", "Bot 2"})
	Game(&game)
}

func Benchmark_Game(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		game := grabble.CreateDefaultGame([]string{"Bot 1", "Bot 2"})
		b.StartTimer()
		Game(&game)
	}
}
