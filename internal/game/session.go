package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	players      map[uuid.UUID]*Player
	playersOrder []uuid.UUID
	snakes       map[uuid.UUID]*objs.Snake
	field        *objs.Field
	stateServer  StateServer
}

func NewSession(players []*Player, stateServer StateServer) *Session {
	session := &Session{
		players:      make(map[uuid.UUID]*Player),
		playersOrder: make([]uuid.UUID, 0),
		snakes:       make(map[uuid.UUID]*objs.Snake),
		stateServer:  stateServer,
		field: objs.NewField(
			common.DefaultFieldHeight,
			common.DefaultFieldWidth,
		),
	}
	for _, player := range players {
		// avoiding equal usernames
		player.username = uniqueUsername(player.username, session.players)
		// players map for fast access
		session.players[player.id] = player
		// players order slice to keep scores order without sorting on client side
		session.playersOrder = append(session.playersOrder, player.id)
		// snake, initial score and apples for each player
		session.snakes[player.id] = session.field.SpawnSnake()
		session.players[player.id].score = common.DefaultScoreOnStart
		for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
			session.field.SpawnApple()
		}
	}
	return session
}

func (s *Session) Run() {
	for {
		s.tick()
	}
}

func (s *Session) ChangePlayersDirection(playerId uuid.UUID, direction common.MoveDirection) {
	if snake := s.snakes[playerId]; snake != nil {
		snake.ChangeDirection(direction)
	} else {
		// todo: log error
	}
}

func (s *Session) tick() {
	future := make(map[uuid.UUID]common.ObjectPosition)
	for playerId, snake := range s.snakes {
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
		if s.field.CheckCellType(position) == objs.CellSnake {
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
			if positionA == s.snakes[playerB].Head().Position() &&
				positionB == s.snakes[playerA].Head().Position() {
				toKill[playerA] = true
				toKill[playerB] = true
			}
		}
	}

	// move
	for playerId, snake := range s.snakes {
		if toKill[playerId] {
			snake.Kill()
			s.players[playerId].score = common.DefaultScoreOnStart
			s.snakes[playerId] = s.field.SpawnSnake()
			continue
		}

		switch s.field.CheckCellType(future[playerId]) {
		case objs.CellApple:
			s.players[playerId].score += common.DefaultScoreIncrease
			snake.Grow()
			s.field.ClearCell(future[playerId])
			s.field.SpawnApple()
			snake.Move()
		case objs.CellEmpty:
			snake.Move()
		default:
		}
	}

	s.stateServer.SendPublicState(s.buildPublicState())
	time.Sleep(100 * time.Millisecond)
}

func (s *Session) buildPublicState() *SessionModel {
	scores := make([]*ScoreModel, 0)
	for _, playerId := range s.playersOrder {
		scores = append(scores, &ScoreModel{
			username: s.players[playerId].username,
			score:    s.players[playerId].score,
		})
	}
	return &SessionModel{
		Scores: scores,
		Field: &objs.FieldModel{
			Height: s.field.Height(),
			Width:  s.field.Width(),
			Cells:  s.field.CellsMap(),
		},
	}
}
