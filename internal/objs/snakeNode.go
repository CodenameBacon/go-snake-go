package objs

import "go-snake-go/internal/common"

type SnakeNode struct {
	next     *SnakeNode
	position common.ObjectPosition
}

func NewSnakeNode(posX, posY int, next *SnakeNode) *SnakeNode {
	return &SnakeNode{
		next:     next,
		position: common.ObjectPosition{X: posX, Y: posY},
	}
}

func (sn *SnakeNode) Position() common.ObjectPosition {
	return sn.position
}

func (sn *SnakeNode) Next() *SnakeNode {
	return sn.next
}
