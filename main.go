package main

import (
	"fmt"
	"stocks/stocks"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	p.Title.Text = "Vector Rotation Visualization"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	vec := stocks.Vector2D{X: 10, Y: 10}
	var points plotter.XYs

	for i := 0; i <= 120; i++ {
		points = append(points, plotter.XY{X: vec.X, Y: vec.Y})
		vec = *stocks.RandomVector2D()
		fmt.Println(vec)
	}

	line, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	p.X.Min = -1
	p.X.Max = 1

	p.Y.Min = -1
	p.Y.Max = 1
	
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "vector_rotation.png"); err != nil {
		panic(err)
	}

	fmt.Println("Vector rotation plot saved as 'vector_rotation.png'")
}
