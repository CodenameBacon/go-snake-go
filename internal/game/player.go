package game

import (
	"go-snake-go/internal/common"

	"github.com/google/uuid"
)

type Player struct {
	Id       uuid.UUID `json:"Id"`
	Username string    `json:"Username"`
	Score    int       `json:"Score"`
}

func NewPlayer(username string) *Player {
	return &Player{
		Id:       uuid.New(),
		Username: username,
		Score:    common.DefaultScoreOnStart,
	}
}
