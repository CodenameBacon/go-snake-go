package game

import (
	"go-snake-go/internal/objs"
)

type ScoreModel struct {
	Username string
	Score    int
}

type SessionModel struct {
	Scores []*ScoreModel    `json:"scores"`
	Field  *objs.FieldModel `json:"field"`
}
