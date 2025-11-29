package objs

type Apple struct {
	posX, posY int
}

func NewApple(posX, posY int) *Apple {
	return &Apple{
		posX: posX,
		posY: posY,
	}
}

func (a *Apple) Position() (int, int) {
	return a.posX, a.posY
}

func (a *Apple) CheckIntersection(snake *Snake) bool {
	if a.Position() == snake.head.Position() {
		return true // intersects with head
	}
	// not checking intersection with other snake body parts
	// cause apple should not be spawned in it
	return false // not intersecting
}
