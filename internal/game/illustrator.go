package game

import (
	"fmt"
	"go-snake-go/internal/common"
)

type SoloGameIllustrator struct{}

func (gi *SoloGameIllustrator) ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func (gi *SoloGameIllustrator) DrawGameField(game *Game) {
	width := game.field.Width() + 2   // + 2 for borders
	height := game.field.Height() + 2 // + 2 for borders

	// создаём поле
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			// стенки по краям
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				grid[i][j] = rune(common.FieldWall[0])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	// рисуем яблоки
	for _, apple := range game.apples {
		posY := apple.Position().Y
		posX := apple.Position().X
		if posY >= 0 && posY < height &&
			posX >= 0 && posX < width {
			grid[posY+1][posX+1] = rune(common.Apple[0]) // + 1 for borders
		}
	}

	// рисуем змею
	node := game.snake.Head()
	for node != nil {
		posY := node.Position().Y
		posX := node.Position().X
		if posY >= 0 && posY < height &&
			posX >= 0 && posX < width {
			grid[posY+1][posX+1] = rune(common.SnakeNode[0]) // + 1 for borders
		}
		node = node.Next()
	}

	// вывод в терминал
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}

	fmt.Printf("Score: %d\n", game.score)
}
