package game

import (
	"go-snake-go/internal/common"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/google/uuid"
)

type ClientManager interface {
	ListenKeyboard()
	sendAction(action PlayerAction)
}

// SoloClientManager - implementation of ClientManager for solo gameplay
type SoloClientManager struct {
	playerId uuid.UUID
	session  *Session
}

func NewSoloClientManager(playerId uuid.UUID, session *Session) *SoloClientManager {
	return &SoloClientManager{playerId: playerId, session: session}
}

func (cm SoloClientManager) ListenKeyboard() {
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
			cm.sendAction(PlayerQuitGame)
			_ = keyboard.Close()
			os.Exit(0)
		case keyboard.KeyArrowUp:
			cm.sendAction(PlayerChangeDirectionUp)
		case keyboard.KeyArrowDown:
			cm.sendAction(PlayerChangeDirectionDown)
		case keyboard.KeyArrowLeft:
			cm.sendAction(PlayerChangeDirectionLeft)
		case keyboard.KeyArrowRight:
			cm.sendAction(PlayerChangeDirectionRight)
		default:
			continue
		}
	}
}

func (cm SoloClientManager) sendAction(action PlayerAction) {
	switch action {
	case PlayerChangeDirectionUp:
		cm.session.ChangePlayersDirection(cm.playerId, common.MoveDirectionUp)
	case PlayerChangeDirectionDown:
		cm.session.ChangePlayersDirection(cm.playerId, common.MoveDirectionDown)
	case PlayerChangeDirectionLeft:
		cm.session.ChangePlayersDirection(cm.playerId, common.MoveDirectionLeft)
	case PlayerChangeDirectionRight:
		cm.session.ChangePlayersDirection(cm.playerId, common.MoveDirectionRight)
	}
}
