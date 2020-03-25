package mock

import (
	"github.com/stretchr/testify/mock"
)

type IntRangedRandomMock struct {
	mock.Mock
}

func (i IntRangedRandomMock) RandomNumber(intRage int) int {
	args := i.Called(intRage)
	return args.Int(0)
}
