package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
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

	for position, cellType := range sessionModel.Field.Cells {
		switch cellType {
		case objs.CellSnake:
			grid[position.Y+1][position.X+1] = common.SnakeNode // + 1 for borders
		case objs.CellApple:
			grid[position.Y+1][position.X+1] = common.Apple // + 1 for borders
		}
	}

	for _, row := range grid {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteByte('\n')
	}

	os.Stdout.Write([]byte(sb.String()))
}
