package solo

import (
	"go-snake-go/internal/common"
	"go-snake-go/internal/objs"
	"time"
)

type Game struct {
	score       int
	field       *objs.Field
	snake       *objs.Snake
	apples      []*objs.Apple
	illustrator *GameIllustrator
}

func NewGame() *Game {
	newGame := &Game{
		score:       0,
		field:       objs.NewField(common.DefaultFieldHeight, common.DefaultFieldWidth),
		illustrator: &GameIllustrator{}, // fixme: implement constructor
	}
	for i := 0; i < common.DefaultTotalApplesOnStart; i++ {
		newGame.SpawnApple()
	}
	newGame.SpawnSnake()
	return newGame
}

func (g *Game) Run() {
	go func() {
		for {
			g.tick()
		}
	}()
	go ListenKeyboard(g)
}

func (g *Game) tick() {
	g.snake.Move()
	g.illustrator.ClearScreen()
	g.illustrator.DrawGameField(g)
	if g.snake.CheckSnakeIntersection(g.snake) {
		g.snake = nil // kills snake
		g.score = 0   // resets score
		g.SpawnSnake()
	} else {
		eatenApple := -1
		for index, apple := range g.apples {
			if g.snake.CheckAppleIntersection(apple) {
				g.score += 100
				g.SpawnApple()
				g.snake.AddTail()
				eatenApple = index
			}
		}
		if eatenApple >= 0 {
			g.apples = append(g.apples[:eatenApple], g.apples[eatenApple+1:]...)
		}
	}
	time.Sleep(100 * time.Millisecond)
}

func (g *Game) SpawnSnake() {
	snake := objs.NewSnake(g.field)

	// fixme: implement no loop variation (based on map probably)
	for func() bool {
		result := false
		for _, ap := range g.apples {
			result = result || snake.CheckAppleIntersection(ap)
		}
		return result
	}() {
		snake = objs.NewSnake(g.field)
	}
	g.snake = snake
}

func (g *Game) SpawnApple() {
	apple := objs.NewApple(g.field)

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
		apple = objs.NewApple(g.field)
	}
	g.apples = append(g.apples, apple)
}
