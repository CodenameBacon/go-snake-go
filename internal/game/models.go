package game

import (
	"go-snake-go/internal/objs"

	"github.com/google/uuid"
)

type ScoreModel struct {
	username string
	score    int
}

type SessionModel struct {
	Scores []*ScoreModel    `json:"scores"`
	Field  *objs.FieldModel `json:"field"`
}

type CurrentPlayerModel struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
