package game

import (
	"fmt"
	"go-snake-go/internal/common"
	"os"
	"strings"
)

func clearScreen() {
	fmt.Fprint(os.Stdout, "\033[2J\033[H")
}

func drawGameField(sessionModel *SessionModel) {
	width := sessionModel.Field.Width + 2   // + 2 for borders
	height := sessionModel.Field.Height + 2 // + 2 for borders

	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				grid[i][j] = common.FieldWall
			} else {
				grid[i][j] = common.Empty
			}
		}
	}

	for _, apple := range sessionModel.Apples {
		posY := apple.Position.Y
		posX := apple.Position.X
		if posY >= 0 && posY < height &&
			posX >= 0 && posX < width {
			grid[posY+1][posX+1] = common.Apple // + 1 for borders
		}
	}

	for _, snake := range sessionModel.Snakes {
		for _, node := range snake.Nodes {
			if node.Position.Y >= 0 && node.Position.Y < height &&
				node.Position.X >= 0 && node.Position.X < width {
				grid[node.Position.Y+1][node.Position.X+1] = common.SnakeNode // + 1 for borders
			}
		}
	}

	for _, row := range grid {
		fmt.Fprintln(os.Stdout, strings.Join(row, ""))
	}

	// todo: draw scores
}
