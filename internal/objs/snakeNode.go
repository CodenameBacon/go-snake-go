package objs

import "go-snake-go/internal/common"

type SnakeNode struct {
	next     *SnakeNode
	position common.ObjectPosition
}

func NewSnakeNode(position common.ObjectPosition, next *SnakeNode) *SnakeNode {
	return &SnakeNode{
		next:     next,
		position: position,
	}
}

func (sn *SnakeNode) Position() common.ObjectPosition {
	return sn.position
}

func (sn *SnakeNode) Next() *SnakeNode {
	return sn.next
}
