package objs

import "go-snake-go/internal/common"

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

func (f *Field) SpawnApple() {
	apple := NewApple(common.GetRandomPosition(f.height, f.width))

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		if f.snake != nil {
			result = f.snake.CheckAppleIntersection(apple)
		}
		for _, ap := range f.apples {
			result = result || ap.CheckAppleIntersection(apple)
		}
		return result
	}() {
		apple = NewApple(common.GetRandomPosition(f.height, f.width))
	}
	f.apples = append(f.apples, apple)
}

func (f *Field) SpawnSnake() {
	snake := NewSnake(common.GetRandomPosition(f.height, f.width))

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, ap := range f.apples {
			result = result || snake.CheckAppleIntersection(ap)
		}
		return result
	}() {
		snake = NewSnake(common.GetRandomPosition(f.height, f.width))
	}
	f.snake = snake
}
