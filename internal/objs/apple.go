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

func (a *Apple) CheckAppleIntersection(apple *Apple) bool {
	return a.Position() == apple.Position()
}
