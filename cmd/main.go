package main

import (
	"go-snake-go/internal/screen"
)

func main() {
	screen.InitTerminal()
	l := screen.NewLauncher()
	l.Run(screen.MenuScreen)
	screen.ResetTerminal()
}
