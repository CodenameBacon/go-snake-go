package common

import "math/rand"

func GetRandomPosition(maxHeight, maxWidth int) (int, int) {
	return rand.Intn(maxWidth - 1), rand.Intn(maxHeight - 1)
}
