package solo

import (
	"go-snake-go/internal/common"
	"os"

	"github.com/eiannone/keyboard"
)

func ListenKeyboard(game *Game) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch key {
		case keyboard.KeyEsc:
			_ = keyboard.Close()
			os.Exit(0)
		case keyboard.KeyArrowUp:
			if game.snake.CurrentDirection() != common.MoveDirectionDown {
				game.snake.ChangeDirection(common.MoveDirectionUp)
			}
		case keyboard.KeyArrowDown:
			if game.snake.CurrentDirection() != common.MoveDirectionUp {
				game.snake.ChangeDirection(common.MoveDirectionDown)
			}
		case keyboard.KeyArrowLeft:
			if game.snake.CurrentDirection() != common.MoveDirectionRight {
				game.snake.ChangeDirection(common.MoveDirectionLeft)
			}
		case keyboard.KeyArrowRight:
			if game.snake.CurrentDirection() != common.MoveDirectionLeft {
				game.snake.ChangeDirection(common.MoveDirectionRight)
			}
		default:
			continue
		}
	}
}
