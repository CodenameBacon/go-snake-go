package game

import (
	"go-snake-go/internal/objs"

	"github.com/google/uuid"
)

type SessionModel struct {
	Scores map[uuid.UUID]int  `json:"scores"`
	Snakes []*objs.SnakeModel `json:"snakes"`
	Apples []*objs.AppleModel `json:"apples"`
	Field  *objs.FieldModel   `json:"field"`
}

type CurrentPlayerModel struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
