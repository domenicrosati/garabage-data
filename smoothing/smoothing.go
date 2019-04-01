package smoothing

type Smoother interface {
	SetWindow(int)
	Smooth([]float64) []float64
}

type NoSmooth struct{}

func (s NoSmooth) SetWindow(int) {}

func (s NoSmooth) Smooth(data []float64) []float64 {
	return data
}
