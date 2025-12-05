package screen

import "github.com/eiannone/keyboard"

type Launcher struct {
	screens []Screen
	keyChan chan keyboard.Key
	run     bool
}

func NewLauncher() *Launcher {
	return &Launcher{
		keyChan: make(chan keyboard.Key, 10),
		screens: make([]Screen, 0),
		run:     false,
	}
}

func (l *Launcher) OpenScreen(screen Screen) {
	l.screens = append(l.screens, screen)
	screen.Open(l)
}

func (l *Launcher) CloseScreen() {
	if len(l.screens) == 0 {
		return
	}
	l.screens[len(l.screens)-1].Close(l)
	l.screens = l.screens[:len(l.screens)-1]
	if len(l.screens) > 0 {
		l.screens[len(l.screens)-1].Open(l)
	} else {
		l.run = false
	}
}

func (l *Launcher) Run(screen Screen) {
	l.screens = append(l.screens, screen)
	l.screens[0].Open(l)
	l.run = true

	// listen keys
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() { _ = keyboard.Close() }()
	go func() {
		for l.run {
			_, key, err := keyboard.GetKey()
			if err != nil {
				return
			}
			l.keyChan <- key
		}
	}()

	// handle keys with screen's logics
	for l.run {
		curr := l.screens[len(l.screens)-1]
		curr.HandleKeys(l)
	}
}
