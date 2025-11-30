package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	players     []*Player
	snakes      map[uuid.UUID]*objs.Snake
	scores      map[uuid.UUID]int
	field       *objs.Field
	stateServer StateServer
}

func NewSession(players []*Player, stateServer StateServer) *Session {
	session := &Session{
		players:     players,
		snakes:      make(map[uuid.UUID]*objs.Snake),
		scores:      make(map[uuid.UUID]int),
		stateServer: stateServer,
		field: objs.NewField(
			common.DefaultFieldHeight,
			common.DefaultFieldWidth,
		),
	}
	// snake, initial score and apples for each player
	for _, player := range players {
		session.snakes[player.id] = session.field.SpawnSnake()
		session.scores[player.id] = common.DefaultScoreOnStart
		for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
			session.field.SpawnApple()
		}
	}
	return session
}

func (g *Session) Run() {
	for {
		g.tick()
	}
}

func (g *Session) ChangePlayersDirection(playerId uuid.UUID, direction common.MoveDirection) {
	if snake := g.snakes[playerId]; snake != nil {
		snake.ChangeDirection(direction)
	} else {
		// todo: log error
	}
}

func (g *Session) tick() {
	future := make(map[uuid.UUID]common.ObjectPosition)
	for playerId, snake := range g.snakes {
		future[playerId] = snake.GetHeadPositionAfterMove()
	}

	toKill := make(map[uuid.UUID]bool)

	// using slice as map value to handle multiple heads intersection
	toSetAsSnake := make(map[common.ObjectPosition][]uuid.UUID)
	for playerId, position := range future {
		toSetAsSnake[position] = append(toSetAsSnake[position], playerId)
	}

	// handling multiple heads intersection
	for _, playerIds := range toSetAsSnake {
		if len(playerIds) > 1 {
			for _, playerId := range playerIds {
				toKill[playerId] = true
			}
		}
	}

	// handling intersection with snake body parts
	for playerId, position := range future {
		if toKill[playerId] {
			continue
		}
		if g.field.CheckCellType(position) == objs.CellSnake {
			toKill[playerId] = true
		}
	}

	// handling head swaps case
	for playerA, positionA := range future {
		if toKill[playerA] {
			continue
		}
		for playerB, positionB := range future {
			if playerA == playerB || toKill[playerB] {
				continue
			}
			// killing both
			if positionA == g.snakes[playerB].Head().Position() &&
				positionB == g.snakes[playerA].Head().Position() {
				toKill[playerA] = true
				toKill[playerB] = true
			}
		}
	}

	// move
	for playerId, snake := range g.snakes {
		if toKill[playerId] {
			snake.Kill()
			g.scores[playerId] = common.DefaultScoreOnStart
			g.snakes[playerId] = g.field.SpawnSnake()
			continue
		}

		switch g.field.CheckCellType(future[playerId]) {
		case objs.CellApple:
			g.scores[playerId] += common.DefaultScoreIncrease
			snake.Grow()
			g.field.ClearCell(future[playerId])
			g.field.SpawnApple()
			snake.Move()
		case objs.CellEmpty:
			snake.Move()
		default:
		}
	}

	time.Sleep(100 * time.Millisecond)
	g.stateServer.SendPublicState(g.buildPublicState())
}

func (g *Session) buildPublicState() *SessionModel {
	return &SessionModel{
		Scores: g.scores,
		Field: &objs.FieldModel{
			Height: g.field.Height(),
			Width:  g.field.Width(),
			Cells:  g.field.CellsMap(),
		},
	}
}
