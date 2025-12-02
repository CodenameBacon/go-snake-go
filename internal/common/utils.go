package common

import (
	"math/rand"
)

func GetRandomPosition(maxHeight, maxWidth int) ObjectPosition {
	return ObjectPosition{rand.Intn(maxWidth - 1), rand.Intn(maxHeight - 1)}
}
