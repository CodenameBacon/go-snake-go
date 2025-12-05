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
		head:      NewSnakeNode(field.GetEmptyPosition(), nil),
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

func (s *Snake) ChangeDirection(direction common.MoveDirection) {
	switch direction {
	case common.MoveDirectionUp:
		if s.CurrentDirection() != common.MoveDirectionDown {
			s.direction = common.MoveDirectionUp
		}
	case common.MoveDirectionDown:
		if s.CurrentDirection() != common.MoveDirectionUp {
			s.direction = common.MoveDirectionDown
		}
	case common.MoveDirectionLeft:
		if s.CurrentDirection() != common.MoveDirectionRight {
			s.direction = common.MoveDirectionLeft
		}
	case common.MoveDirectionRight:
		if s.CurrentDirection() != common.MoveDirectionLeft {
			s.direction = common.MoveDirectionRight
		}
	}
}

// Grow - adds new SnakeNode as tail which will be placed as head on the next tick
func (s *Snake) Grow() {
	head := s.head
	for head.next != nil {
		head = head.next
	}
	head.next = NewSnakeNode(head.position, nil)
}

func (s *Snake) Move() {
	newHeadPos := s.GetHeadPositionAfterMove()

	// handling move of snake with len = 1
	if s.head.next == nil {
		s.field.ClearCell(s.head.position)
		s.head.position = newHeadPos
		s.field.SetCellType(newHeadPos, CellSnakeHead)
		return
	}

	// handling move of snake with len > 1
	preTail := s.head
	for preTail.next != nil && preTail.next.next != nil {
		preTail = preTail.next
	}
	if preTail.next != nil {
		s.field.ClearCell(preTail.next.position)
		preTail.next = nil
		s.field.SetCellType(preTail.position, CellSnakeTail)
	}

	s.head = NewSnakeNode(newHeadPos, s.head)
	s.field.SetCellType(newHeadPos, CellSnakeHead)
	s.field.SetCellType(s.head.next.position, CellSnakeNode)
}

func (s *Snake) Kill() {
	node := s.head
	for node != nil {
		s.field.ClearCell(node.position)
		node = node.next
	}
	s.head = nil
}

func (s *Snake) GetHeadPositionAfterMove() common.ObjectPosition {
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
		panic(fmt.Sprintf("Impossible Move for perform: %v", s.direction))
	}
	return position
}
