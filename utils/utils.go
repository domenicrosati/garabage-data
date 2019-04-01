package utils

import (
	"fmt"
	"os"

	chart "github.com/wcharczuk/go-chart"
	"gonum.org/v1/gonum/mat"
)

func Ones(window, length int) *mat.Dense {
	data := []float64{}
	for _ = range make([]int, length) {
		data = append(data, 1.0)
	}
	return mat.NewDense(window, len(data)/window, data)
}

func Arange(l int) []float64 {
	data := []float64{}
	for t := range make([]int, l) {
		data = append(data, float64(t))
	}
	return data
}

func fill(data, origData []float64) []float64 {
	interpolated := []float64{}
	padding := len(origData) / len(data)
	for i := range make([]int, len(origData)) {
		sample := data[i/padding]
		interpolated = append(interpolated, sample)
	}
	return interpolated
}

func CreateChart(data []float64, origData []float64, filename string) error {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: Arange(len(origData)),
				YValues: fill(data, origData),
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.ColorOrange,
				},
				XValues: Arange(len(origData)),
				YValues: origData,
			},
		},
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = graph.Render(chart.PNG, f)
	if err != nil {
		return err
	}
	return nil
}

func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
