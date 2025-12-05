package screen

type BaseScreenAction struct {
	title     string
	toExecute func(l *Launcher)
}

func (sa *BaseScreenAction) Text() string {
	return sa.title
}

func (sa *BaseScreenAction) Execute(l *Launcher) {
	sa.toExecute(l)
}
