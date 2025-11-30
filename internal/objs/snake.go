package objs

import (
	"fmt"
	"go-snake-go/internal/common"
)

type Snake struct {
	field     *Field
	head      *SnakeNode
	direction common.MoveDirection
}

func NewSnake(field *Field) *Snake {
	return &Snake{
		head:      NewSnakeNode(common.GetRandomPosition(field.height, field.width), nil),
		direction: common.DefaultMoveDirectionOnStart,
		field:     field,
	}
}

func (s *Snake) Head() *SnakeNode {
	return s.head
}

func (s *Snake) CurrentDirection() common.MoveDirection {
	return s.direction
}

func (s *Snake) ChangeDirection(dir common.MoveDirection) {
	s.direction = dir
}

// Grow - adds new SnakeNode as tail which will be placed as head on the next tick
func (s *Snake) Grow() {
	head := s.head
	for head.next != nil {
		head = head.next
	}
	head.next = NewSnakeNode(head.position, nil)
}

// Move - places tail node as new head on the position after move
func (s *Snake) Move() {
	headPosition := s.getHeadPositionAfterMove()
	preTail := s.head
	for preTail.next != nil && preTail.next.next != nil {
		preTail = preTail.next
	}
	oldHead := s.head
	// if snakeLength > 1
	if preTail.next != nil {
		s.head = preTail.next
	}
	s.head.position = headPosition
	s.head.next = oldHead
	preTail.next = nil
}

func (s *Snake) CheckAppleIntersection(apple *Apple) bool {
	node := s.head
	for node != nil {
		if node.position == apple.position {
			return true
		}
		node = node.next
	}
	return false
}

func (s *Snake) CheckSnakeIntersection(snake *Snake) bool {
	if &s.head != &snake.head && s.head.position == snake.head.position {
		return true
	}
	node := s.head
	for node.next != nil {
		node = node.next
		if node.position == snake.head.position {
			return true
		}
	}
	return false
}

func (s *Snake) getHeadPositionAfterMove() common.ObjectPosition {
	position := s.head.position
	actions := map[common.MoveDirection]func(){
		common.MoveDirectionUp: func() {
			position.Y -= 1
			if position.Y < 0 {
				position.Y = s.field.height - 1
			}
		},
		common.MoveDirectionDown: func() {
			position.Y += 1
			if position.Y > s.field.height-1 {
				position.Y -= s.field.height
			}
		},
		common.MoveDirectionLeft: func() {
			position.X -= 1
			if position.X < 0 {
				position.X = s.field.width - 1
			}
		},
		common.MoveDirectionRight: func() {
			position.X += 1
			if position.X > s.field.width-1 {
				position.X -= s.field.width
			}
		},
	}
	if actions[s.direction] != nil {
		actions[s.direction]()
	} else {
		panic(fmt.Sprintf("Impossible Move for perform: %s", s.direction))
	}
	return position
}
