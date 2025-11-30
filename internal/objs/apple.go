package objs

import "go-snake-go/internal/common"

type Apple struct {
	position common.ObjectPosition
}

func NewApple(field *Field) *Apple {
	return &Apple{
		position: common.GetRandomPosition(field.height, field.width),
	}
}

func (a *Apple) Position() common.ObjectPosition {
	return a.position
}

func (a *Apple) CheckAppleIntersection(apple *Apple) bool {
	return a.position == apple.position
}
