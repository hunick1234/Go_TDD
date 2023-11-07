package game_test

import (
	"snake/game"
	"testing"
)

func TestMove(t *testing.T) {
	g := &game.GameObject{0, 0}
	initialY := g.Y
	initialX := g.X

	t.Run("MoveDown", func(t *testing.T) {
		g.MoveDown()
		if g.Y != initialY-1 {
			t.Errorf("MoveDown failed, expected %d, got %d", initialY-1, g.Y)
		}
	})

	t.Run("MoveUp", func(t *testing.T) {
		g.MoveUp()
		if g.Y != initialY {
			t.Errorf("MoveUp failed, expected %d, got %d", initialY, g.Y)
		}
	})

	t.Run("MoveLeft", func(t *testing.T) {
		g.MoveLeft()
		if g.X != initialX-1 {
			t.Errorf("MoveLeft failed, expected %d, got %d", initialX-1, g.X)
		}
	})

	t.Run("MoveRight", func(t *testing.T) {
		g.MoveRight()
		if g.X != initialX {
			t.Errorf("MoveRight failed, expected %d, got %d", initialX, g.X)
		}
	})

}
