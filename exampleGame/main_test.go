package main

import (
	"flag"
	"testing"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func TestGame(t *testing.T) {
	Game()
}

func Benchmark_Game(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Game()
	}
}
