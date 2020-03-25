package random

import (
	"github.com/juanimoli/piccadilly/pkg/data/usecase"
	"math/rand"
)

type intRangedRandom struct {
	rand *rand.Rand
}

func CreateIntRangedRandom(rand *rand.Rand) usecase.IntRagedRandom {
	return intRangedRandom{
		rand: rand,
	}
}

func (i intRangedRandom) RandomNumber(intRage int) int {
	return i.rand.Intn(intRage)
}
