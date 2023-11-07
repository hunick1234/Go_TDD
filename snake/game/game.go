package game

import (
	"fmt"
	_ "os"
)

func RenderGame() {
	fmt.Println("RenderGame")
}

func InitGame() [10][10]int {
	var gameMap [10][10]int
	return gameMap
}
