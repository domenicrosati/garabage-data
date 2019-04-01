package smoothing

import (
	"github.com/domenicrosati/garbage-data/utils"
	"gonum.org/v1/gonum/mat"
)

type MovingAverage struct {
	Window int
}

func (m MovingAverage) SetWindow(w int) {
	m.Window = w
}

//  1'X(1'1)-1
func (m MovingAverage) Smooth(data []float64) []float64 {
	X := mat.NewDense(len(data)/m.Window, m.Window, data)
	one := utils.Ones(1, m.Window)

	weights := mat.NewDense(1, 1, nil)
	weights.Mul(one, one.T())
	weights.Inverse(weights)

	sum := mat.NewDense(1, len(data)/m.Window, nil)
	sum.Product(one, X.T())

	Y := mat.NewDense(1, len(data)/m.Window, nil)
	Y.Product(weights, sum)
	return Y.RawRowView(0)
}
