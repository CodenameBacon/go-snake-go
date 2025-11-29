package game

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
	"time"
)

type Game struct {
	score  int
	field  *objs.Field
	snake  *objs.Snake
	apples []*objs.Apple
	// todo: implement illustrator
	// todo: implement keyboard listener
}

func NewGame() *Game {
	newGame := &Game{
		score: 0,
		field: objs.NewField(common.DefaultFieldHeight, common.DefaultFieldWidth),
	}
	for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
		newGame.SpawnApple()
	}
	newGame.SpawnSnake()
	return newGame
}

func (g *Game) Run() {
	go func() {
		for g.snake != nil {
			g.tick()
		}
	}()
}

func (g *Game) tick() {
	if g.snake == nil {
		g.SpawnSnake()
	}
	g.snake.Move()
	// todo: illustrate
	if g.snake.CheckSnakeIntersection(g.snake) {
		g.snake = nil // kills snake
		g.score = 0   // resets score
	} else {
		for _, apple := range g.apples {
			if g.snake.CheckAppleIntersection(apple) {
				g.score += 100
			}
		}
	}
	time.Sleep(250 * time.Millisecond)
}

func (g *Game) SpawnSnake() {
	snake := objs.NewSnake(*g.field)

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, ap := range g.apples {
			result = result || snake.CheckAppleIntersection(ap)
		}
		return result
	}() {
		snake = objs.NewSnake(*g.field)
	}
	g.snake = snake
}

func (g *Game) SpawnApple() {
	apple := objs.NewApple(common.GetRandomPosition(g.field.Height(), g.field.Width()))

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		if g.snake != nil {
			result = g.snake.CheckAppleIntersection(apple)
		}
		for _, ap := range g.apples {
			result = result || ap.CheckAppleIntersection(apple)
		}
		return result
	}() {
		apple = objs.NewApple(
			common.GetRandomPosition(
				g.field.Height(),
				g.field.Width(),
			),
		)
	}
	g.apples = append(g.apples, apple)
}
