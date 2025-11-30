package game

import (
	"github.com/google/uuid"
)

type PlayerAction int

const (
	PlayerChangeDirectionUp    PlayerAction = 1
	PlayerChangeDirectionDown  PlayerAction = 2
	PlayerChangeDirectionLeft  PlayerAction = 3
	PlayerChangeDirectionRight PlayerAction = 4
	PlayerQuitGame             PlayerAction = 5
)

type Player struct {
	id       uuid.UUID `json:"id"`
	username string    `json:"username"`
}

func NewPlayer(username string) *Player {
	return &Player{
		id:       uuid.New(),
		username: username,
	}
}

func (p *Player) Id() uuid.UUID {
	return p.id
}
