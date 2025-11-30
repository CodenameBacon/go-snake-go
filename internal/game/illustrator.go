package game

import (
	"go-snake-go/internal/common"
	"os"
	"strings"
)

func InitTerminal() {
	os.Stdout.Write(common.HideCursor)
	os.Stdout.Write(common.CursorHome)
	os.Stdout.Write(common.ClearBelow)
}

func ResetTerminal() {
	os.Stdout.Write(common.ShowCursor)
}

func drawGameField(sessionModel *SessionModel) {
	width := sessionModel.Field.Width + 2
	height := sessionModel.Field.Height + 2
	os.Stdout.Write(common.CursorHome)

	sb := strings.Builder{}
	sb.Grow(width * height * 2)

	grid := make([][]string, height)
	for i := range grid {
		row := make([]string, width)
		for j := range row {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				row[j] = common.FieldWall
			} else {
				row[j] = common.Empty
			}
		}
		grid[i] = row
	}

	for _, apple := range sessionModel.Apples {
		y := apple.Position.Y + 1
		x := apple.Position.X + 1
		if y >= 1 && y < height-1 && x >= 1 && x < width-1 {
			grid[y][x] = common.Apple
		}
	}

	for _, snake := range sessionModel.Snakes {
		for _, node := range snake.Nodes {
			y := node.Position.Y + 1
			x := node.Position.X + 1
			if y >= 1 && y < height-1 && x >= 1 && x < width-1 {
				grid[y][x] = common.SnakeNode
			}
		}
	}

	for _, row := range grid {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteByte('\n')
	}

	os.Stdout.Write([]byte(sb.String()))
}
