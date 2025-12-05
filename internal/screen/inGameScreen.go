package screen

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/game"
	"os"

	"github.com/eiannone/keyboard"
)

type SoloInGameScreen struct {
	title            string
	player           *game.Player
	sessionStateChan chan *game.SessionModel
	inputActionChan  chan *game.InputAction
}

func (s *SoloInGameScreen) Open(_ *Launcher) {
	s.sessionStateChan = make(chan *game.SessionModel)

	player := game.NewPlayer("you")
	session := game.NewSession(
		[]*game.Player{player},
		&game.SoloStateServer{StateChan: s.sessionStateChan},
	)
	s.inputActionChan = session.InputActionChan
	s.player = player
	go session.Run()
	go s.handleSessionState()
}

func (s *SoloInGameScreen) Close(_ *Launcher) {
	panic("Implement me!")
	// todo:
	// 	stop session
	//  stop handleSessionState method
}

func (s *SoloInGameScreen) handleSessionState() {
	for {
		select {
		case state := <-s.sessionStateChan:
			os.Stdout.Write(common.CursorHome)
			os.Stdout.Write([]byte("go-snake-go: " + s.title + "\n"))
			os.Stdout.Write([]byte("\n"))
			drawSessionState(state)
		}
	}
}

func (s *SoloInGameScreen) HandleKeys(l *Launcher) {
	select {
	case key := <-l.keyChan:
		switch key {
		case keyboard.KeyEsc:
			panic("Implement me!")
		case keyboard.KeyArrowUp:
			s.inputActionChan <- &game.InputAction{
				PlayerId: s.player.Id(),
				ActionId: game.InputChangeDirUp,
			}
		case keyboard.KeyArrowDown:
			s.inputActionChan <- &game.InputAction{
				PlayerId: s.player.Id(),
				ActionId: game.InputChangeDirDown,
			}
		case keyboard.KeyArrowLeft:
			s.inputActionChan <- &game.InputAction{
				PlayerId: s.player.Id(),
				ActionId: game.InputChangeDirLeft,
			}
		case keyboard.KeyArrowRight:
			s.inputActionChan <- &game.InputAction{
				PlayerId: s.player.Id(),
				ActionId: game.InputChangeDirRight,
			}
		default:
		}
	default:
	}
}
