package game

import (
	"testing"
)

func TestInitGame(t *testing.T) {
	game := InitGame()
	expected := [10][10]int{}
	if len(game) != len(expected) {
		t.Errorf("InitGame failed, expected %d, got %d", len(expected), len(game))
	}
	if len(game[0]) != len(expected[0]) {
		t.Errorf("InitGame failed, expected %d, got %d", len(expected[0]), len(game[0]))
	}
}
