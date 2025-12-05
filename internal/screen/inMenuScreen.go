package screen

import (
	"go-snake-go/internal/common"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

type InMenuScreen struct {
	title   string
	actions []ScreenAction
	cursor  int
	level   int
}

func (s *InMenuScreen) clearTerminal() {
	os.Stdout.Write(common.CursorHome)
	os.Stdout.Write(common.ClearBelow)
}

func (s *InMenuScreen) draw() {
	s.clearTerminal()
	os.Stdout.Write([]byte("go-snake-go: " + s.title + "\n"))
	os.Stdout.Write([]byte("\n"))
	for index, action := range s.actions {
		os.Stdout.Write([]byte(strings.Repeat(">", s.level) + " " + action.Text()))
		if index == s.cursor {
			os.Stdout.Write([]byte(" <--"))
		}
		os.Stdout.Write([]byte("\n"))
	}
}

func (s *InMenuScreen) Open(l *Launcher) {
	s.level = len(l.screens)
	s.draw()
}

func (s *InMenuScreen) Close(_ *Launcher) {
	// no action cause nothing to close
}

func (s *InMenuScreen) HandleKeys(l *Launcher) {
	select {
	case key := <-l.keyChan:
		switch key {
		case keyboard.KeyEnter:
			s.actions[s.cursor].Execute(l)
		case keyboard.KeyEsc:
			l.CloseScreen()
		case keyboard.KeyArrowUp:
			if s.cursor > 0 {
				s.cursor--
			}
			s.draw()
		case keyboard.KeyArrowDown:
			if s.cursor < len(s.actions)-1 {
				s.cursor++
			}
			s.draw()
		default:
		}
	default:
	}
}
