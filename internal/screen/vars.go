package screen

// screens
var (
	MenuScreen = &InMenuScreen{
		title:   ".menu",
		actions: []ScreenAction{openSelectModeScreenAction, exitScreenAction},
	}
	SelectModeScreen = &InMenuScreen{
		title:   ".selectmode",
		actions: []ScreenAction{startSoloGameAction, backToPrevScreenAction},
	}
	InSoloGameScreen = &SoloInGameScreen{title: ".insologame"}
)

// screen actions
var (
	openSelectModeScreenAction ScreenAction = &BaseScreenAction{
		title: "play",
		toExecute: func(l *Launcher) {
			l.OpenScreen(SelectModeScreen)
		},
	}
	startSoloGameAction ScreenAction = &BaseScreenAction{
		title: "solo",
		toExecute: func(l *Launcher) {
			l.OpenScreen(InSoloGameScreen)
		},
	}
	backToPrevScreenAction ScreenAction = &BaseScreenAction{
		title: "back",
		toExecute: func(l *Launcher) {
			l.CloseScreen()
		},
	}
	exitScreenAction ScreenAction = &BaseScreenAction{
		title: "exit",
		toExecute: func(l *Launcher) {
			l.CloseScreen()
		},
	}
)
