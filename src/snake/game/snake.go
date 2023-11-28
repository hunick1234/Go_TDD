package game

type snake struct {
	Pre  *snake
	Back *snake
}

type GameObject struct {
	X, Y int
}

type Mover interface {
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
}

func (s *snake) DeleteEnd() {
	if s.Back != nil {
		s.Back.Pre = nil
	}
}

func (s *snake) AddHead(x, y int) {

}

func (s *GameObject) MoveUp() {
	s.Y++
}

func (s *GameObject) MoveDown() {
	s.Y--
}

func (s *GameObject) MoveLeft() {
	s.X--
}

func (s *GameObject) MoveRight() {
	s.X++
}
