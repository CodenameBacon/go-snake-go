package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
)

type SoloGame struct {
	score int
	field *objs.Field
	// todo: implement illustrator
	// todo: implement keyboard listener
}

func NewSoloGame() *SoloGame {
	field := objs.NewField(
		common.DefaultFieldHeight,
		common.DefaultFieldWidth,
	)
	for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
		field.SpawnApple()
	}
	field.SpawnSnake()
	return &SoloGame{
		score: 0,
		field: field,
	}
}
