package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
	"time"

	"github.com/google/uuid"
)

type GameSession struct {
	players []*Player
	snakes  map[uuid.UUID]*objs.Snake
	scores  map[uuid.UUID]int
	apples  []*objs.Apple
	field   *objs.Field
}

func NewGame(players []*Player) *GameSession {
	game := &GameSession{
		players: players,
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

func (g *GameSession) Run() {
	go func() {
		for {
			g.tick()
		}
	}()
	select {}
}

func (g *GameSession) ChangePlayersDirection(playerId uuid.UUID, direction common.MoveDirection) {
	if snake := g.snakes[playerId]; snake != nil {
		snake.ChangeDirection(direction)
	} else {
		// todo: log error
	}
}

func (g *GameSession) tick() {
	for _, snake := range g.snakes {
		snake.Move()
	}

	for playerId, snake := range g.snakes {
		if snake.CheckSnakeIntersection(snake) {
			snake = nil // kills snake
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
}

func (g *GameSession) spawnSnake(playerId uuid.UUID) {
	snake := objs.NewSnake(g.field)

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, snk := range g.snakes {
			result = snk.CheckSnakeIntersection(snk)
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

func (g *GameSession) spawnApple() {
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
