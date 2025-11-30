package game

import (
	"go-snake-go/internal/common"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/google/uuid"
)

type ClientManager interface {
	ListenKeyboard()
	sendAction()
}

// SoloClientManager - implementation of ClientManager for solo gameplay
type SoloClientManager struct {
	playerId uuid.UUID
	session  *Session

	actionsChan chan PlayerAction
}

func NewSoloClientManager(playerId uuid.UUID, session *Session) *SoloClientManager {
	cm := &SoloClientManager{
		playerId:    playerId,
		session:     session,
		actionsChan: make(chan PlayerAction),
	}
	go cm.sendAction()
	return cm
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
			cm.actionsChan <- PlayerQuitGame
			_ = keyboard.Close()
			os.Exit(0)
		case keyboard.KeyArrowUp:
			cm.actionsChan <- PlayerChangeDirectionUp
		case keyboard.KeyArrowDown:
			cm.actionsChan <- PlayerChangeDirectionDown
		case keyboard.KeyArrowLeft:
			cm.actionsChan <- PlayerChangeDirectionLeft
		case keyboard.KeyArrowRight:
			cm.actionsChan <- PlayerChangeDirectionRight
		default:
			continue
		}
	}
}

func (cm SoloClientManager) sendAction() {
	for action := range cm.actionsChan {
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
}
