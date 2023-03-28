package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Shape(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test circle", func(t *testing.T) {
		calculator := NewMockAreaCalculator(ctrl)

		circle := Shape{
			r:              p(10),
			areaCalculator: calculator,
		}

		expectedArea := 3.1415 * 10 * 10
		calculator.EXPECT().GetArea().Return(expectedArea).Times(1)

		square := circle.Square()
		assert.Equal(t, square, expectedArea)

	})
}

func p(v float64) *float64 {
	return &v
}
