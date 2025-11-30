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
	headPosX, headPosY := common.GetRandomPosition(field.Height(), field.Width())
	return &Snake{
		head:      NewSnakeNode(headPosX, headPosY, nil),
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

func (s *Snake) AddTail() {
	head := s.head
	for head.next != nil {
		head = head.next
	}
	// dummy node which will be removed on next tick
	head.next = NewSnakeNode(head.position.X, head.position.Y, nil)
}

func (s *Snake) Move() {
	newPosition := s.getHeadPositionAfterMove()
	newHead := NewSnakeNode(newPosition.X, newPosition.Y, s.head)
	s.head = newHead
	s.removeTail() // fixme: should not be used if snake ate an apple on this Move
}

func (s *Snake) CheckAppleIntersection(apple *Apple) bool {
	node := s.head
	for node != nil {
		if node.position == apple.Position() {
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

// removeTail - removes the last SnakeNode in the ll. Used in Move method.
func (s *Snake) removeTail() {
	head := s.head
	for head.next != nil && head.next.next != nil {
		head = head.next
	}
	head.next = nil
}

func (s *Snake) getHeadPositionAfterMove() common.ObjectPosition {
	position := s.head.position
	actions := map[common.MoveDirection]func(){
		common.MoveDirectionUp: func() {
			position.Y -= 1
			if position.Y < 0 {
				position.Y = s.field.Height() - 1
			}
		},
		common.MoveDirectionDown: func() {
			position.Y += 1
			if position.Y > s.field.Height()-1 {
				position.Y -= s.field.Height()
			}
		},
		common.MoveDirectionLeft: func() {
			position.X -= 1
			if position.X < 0 {
				position.X = s.field.Width() - 1
			}
		},
		common.MoveDirectionRight: func() {
			position.X += 1
			if position.X > s.field.Width()-1 {
				position.X -= s.field.Width()
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
