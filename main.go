package main

import (
	"github.com/domenicrosati/garbage-data/generation"
	"github.com/domenicrosati/garbage-data/smoothing"
	"github.com/domenicrosati/garbage-data/utils"
)

func main() {
	// Generate a noisy workload based on sin wave.
	n := generation.NewNoisySin(360, 360, 30)
	data, _ := n.GenerateWorkload(360)

	utils.CreateChart(data, data, "./data.png")

	// Create a smoother for smoothing data.
	// That is, transforming dirty data => true signal.
	var smoother smoothing.Smoother
	smoother = smoothing.MovingAverage{Window: 10}
	smoothData := smoother.Smooth(data)
	utils.CreateChart(smoothData, data, "./ma_data.png")
}
