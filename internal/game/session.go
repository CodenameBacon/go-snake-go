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
	apples      []*objs.Apple
	field       *objs.Field
	stateServer StateServer
}

func NewSession(players []*Player, stateServer StateServer) *Session {
	game := &Session{
		players:     players,
		snakes:      make(map[uuid.UUID]*objs.Snake),
		scores:      make(map[uuid.UUID]int),
		apples:      make([]*objs.Apple, 0),
		stateServer: stateServer,
		field: objs.NewField(
			common.DefaultFieldHeight,
			common.DefaultFieldWidth,
		),
	}
	// snake, initial score and apples for each player
	for _, player := range players {
		game.spawnSnake(player.id)
		game.scores[player.id] = common.DefaultScoreOnStart
		for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
			game.spawnApple()
		}
	}
	return game
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
	for _, snake := range g.snakes {
		snake.Move()
	}

	for playerId, snake := range g.snakes {
		if snake.CheckSnakeIntersection(snake) {
			g.scores[playerId] = common.DefaultScoreOnStart
			g.spawnSnake(playerId)
		} else {
			eatenApple := -1
			for index, apple := range g.apples {
				if snake.CheckAppleIntersection(apple) {
					g.scores[playerId] += common.DefaultScoreIncrease
					g.spawnApple()
					snake.Grow()
					eatenApple = index
				}
			}
			if eatenApple >= 0 {
				g.apples = append(g.apples[:eatenApple], g.apples[eatenApple+1:]...)
			}
		}
	}
	time.Sleep(100 * time.Millisecond)
	g.stateServer.SendPublicState(g.buildPublicState())
}

func (g *Session) buildPublicState() *SessionModel {
	var apples []*objs.AppleModel
	var snakes []*objs.SnakeModel
	for _, apple := range g.apples {
		apples = append(apples, &objs.AppleModel{Position: apple.Position()})
	}
	for _, snake := range g.snakes {
		node := snake.Head()

		var nodes []*objs.SnakeNodeModel
		for node != nil {
			nodes = append(nodes, &objs.SnakeNodeModel{Position: node.Position()})
			node = node.Next()
		}
		snakes = append(snakes, &objs.SnakeModel{Nodes: nodes})
	}
	data := &SessionModel{
		Scores: g.scores,
		Apples: apples,
		Snakes: snakes,
		Field: &objs.FieldModel{
			Height: g.field.Height(),
			Width:  g.field.Width(),
		},
	}
	return data
}

func (g *Session) spawnSnake(playerId uuid.UUID) {
	snake := objs.NewSnake(g.field)

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, snk := range g.snakes {
			result = snake.CheckSnakeIntersection(snk)
		}
		for _, ap := range g.apples {
			result = result || snake.CheckAppleIntersection(ap)
		}
		return result
	}() {
		snake = objs.NewSnake(g.field)
	}
	g.snakes[playerId] = snake
}

func (g *Session) spawnApple() {
	apple := objs.NewApple(g.field)

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, snk := range g.snakes {
			result = snk.CheckAppleIntersection(apple)
		}
		for _, ap := range g.apples {
			result = result || ap.CheckAppleIntersection(apple)
		}
		return result
	}() {
		apple = objs.NewApple(g.field)
	}
	g.apples = append(g.apples, apple)
}
