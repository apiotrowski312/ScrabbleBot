package main

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func TestGame(t *testing.T) {
	game := grabble.CreateDefaultGame([][2]string{{"Bot 1", "default"}, {"Bot 2", "default"}})
	Game(&game)
}

func Benchmark_Game(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		game := grabble.CreateDefaultGame([][2]string{{"Bot 1", "default"}, {"Bot 2", "default"}})
		b.StartTimer()
		Game(&game)
	}
}
