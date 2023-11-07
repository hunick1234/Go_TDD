package game_test

import (
	"snake/game"
	"testing"
)

func TestMoveUp(t *testing.T) {
	g := &game.GameObject{0, 0}
	initialY := g.Y
	g.MoveUp()
	if g.Y != initialY-1 {
		t.Errorf("MoveUp failed, expected %d, got %d", initialY-1, g.Y)
	}
}
