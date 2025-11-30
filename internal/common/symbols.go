package common

// symbols used to draw field
const (
	SnakeNode string = " @"
	Apple     string = " $"
	FieldWall string = " #"
	Empty     string = "  "
)

var (
	HideCursor = []byte("\033[?25l")
	ShowCursor = []byte("\033[?25h")
	CursorHome = []byte("\033[H")
	ClearBelow = []byte("\033[J")
)
