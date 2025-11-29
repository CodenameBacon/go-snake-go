package main

import (
	"go-snake-go/internal/game/solo"
)

func main() {
	game := solo.NewGame()
	game.Run()
	select {}
}
