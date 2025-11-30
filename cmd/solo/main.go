package main

import (
	"go-snake-go/internal/game"
)

func main() {
	player := game.NewPlayer("you")
	session := game.NewSession(
		append(make([]*game.Player, 0), player),
		&game.SoloStateServer{},
	)
	clientManager := game.NewSoloClientManager(player.Id(), session)
	game.InitTerminal()
	go session.Run()
	go clientManager.ListenKeyboard()
	select {}
}
