package generation

import (
	"math"
	"time"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

type generator interface {
	GenerateWorkload(int64) []float64
}

type NoisySin struct {
	Noise, Amplitude, Period float64
}

func NewNoisySin(noise, amplitude, period float64) *NoisySin {
	return &NoisySin{
		Noise:     noise,
		Amplitude: amplitude,
		Period:    period,
	}
}

func (n *NoisySin) GenerateWorkload(length int64) ([]float64, []float64) {
	samples, actuals := []float64{}, []float64{}
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	noisemaker := distuv.Normal{
		Mu:    n.Noise,
		Sigma: n.Noise,
		Src:   src,
	}
	for t := range make([]int64, length) {
		sample := n.Amplitude*math.Cos(math.Pi/(n.Period*2.0)*float64(t)) + (n.Amplitude + 1.0)
		noise := sample + noisemaker.Rand()
		actuals = append(actuals, sample)
		samples = append(samples, noise)
	}
	return samples, actuals
}
