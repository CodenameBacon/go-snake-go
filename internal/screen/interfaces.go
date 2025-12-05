package screen

type Screen interface {
	Open(l *Launcher)
	HandleKeys(l *Launcher)
	Close(l *Launcher)
}

type ScreenAction interface {
	Text() string
	Execute(l *Launcher)
}
