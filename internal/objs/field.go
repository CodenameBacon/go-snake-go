package objs

type Field struct {
	height, width int
	snake         *Snake
	apples        []*Apple
}

func NewField(height, width int) *Field {
	return &Field{
		height: height,
		width:  width,
	}
}

func (f *Field) Height() int {
	return f.height
}

func (f *Field) Width() int {
	return f.width
}
