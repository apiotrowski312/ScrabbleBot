package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	Game()
}

func Benchmark_Game(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		Game()
		b.StartTimer()
	}
}
