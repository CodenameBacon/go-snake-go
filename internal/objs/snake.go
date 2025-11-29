package objs

import (
	"fmt"
	"go-snake-go/internal/common"
)

type SnakeNode struct {
	posX, posY int
	next       *SnakeNode
}

func NewSnakeNode(posX, posY int) *SnakeNode {
	return &SnakeNode{
		posX: posX,
		posY: posY,
		next: nil,
	}
}

func (sn *SnakeNode) Position() common.ObjectPosition {
	return common.ObjectPosition{X: sn.posX, Y: sn.posY}
}

type Snake struct {
	head      *SnakeNode
	currDir   common.MoveDirection
	gameField Field
}

func NewSnake(field Field) *Snake {
	return &Snake{
		head:      NewSnakeNode(common.GetRandomPosition(field.Height(), field.Width())),
		currDir:   common.DefaultMoveDirectionOnStart,
		gameField: field,
	}
}

// removeTail - removes the last SnakeNode in the ll. Used in Move method.
func (s *Snake) removeTail() {
	head := s.head
	for head.next != nil && head.next.next != nil {
		head = head.next
	}
	head.next = nil
}

func (s *Snake) insertHeadUp() {
	prevHead := s.head
	newPosY := prevHead.posY - 1
	if newPosY < 0 {
		newPosY = s.gameField.Height() - 1
	}
	newHead := &SnakeNode{
		posX: prevHead.posX,
		posY: newPosY,
		next: prevHead,
	}
	s.head = newHead
}

func (s *Snake) insertHeadDown() {
	prevHead := s.head
	newPosY := prevHead.posY + 1
	if newPosY > s.gameField.Height()-1 {
		newPosY = newPosY - s.gameField.Height()
	}
	newHead := &SnakeNode{
		posX: prevHead.posX,
		posY: newPosY,
		next: prevHead,
	}
	s.head = newHead
}

func (s *Snake) insertHeadLeft() {
	prevHead := s.head
	newPosX := prevHead.posX - 1
	if newPosX < 0 {
		newPosX = s.gameField.Width() - 1
	}
	newHead := &SnakeNode{
		posX: newPosX,
		posY: prevHead.posY,
		next: prevHead,
	}
	s.head = newHead
}

func (s *Snake) insertHeadRight() {
	prevHead := s.head
	newPosX := prevHead.posX + 1
	if newPosX > s.gameField.Width()-1 {
		newPosX = newPosX - s.gameField.Width()
	}
	newHead := &SnakeNode{
		posX: newPosX,
		posY: prevHead.posY,
		next: prevHead,
	}
	s.head = newHead
}

func (s *Snake) Move() {
	methodsToCall := map[common.MoveDirection]func(){
		common.MoveDirectionUp:    s.insertHeadUp,
		common.MoveDirectionDown:  s.insertHeadDown,
		common.MoveDirectionLeft:  s.insertHeadLeft,
		common.MoveDirectionRight: s.insertHeadRight,
	}
	if methodsToCall[s.currDir] != nil {
		methodsToCall[s.currDir]()
	} else {
		panic(fmt.Sprintf("Impossible Move for perform: %s", s.currDir))
	}
	s.removeTail() // fixme: should not be used if snake ate an apple on this Move
}

func (s *Snake) CheckAppleIntersection(apple *Apple) bool {
	if s.head.Position() == apple.Position() {
		return true // intersects with head
	}
	node := s.head
	for node.next != nil {
		if node.Position() == apple.Position() {
			return true // intersects with other nodes of snake
		}
	}
	return false // not intersecting
}

func (s *Snake) CheckSnakeIntersection(snake *Snake) bool {
	if &s.head != &snake.head && s.head.Position() == snake.head.Position() {
		return true // intersects with head
	}
	node := s.head
	for node.next != nil {
		if node.Position() == snake.head.Position() {
			return true // intersects with other nodes of snake
		}
	}
	return false // not intersecting
}
