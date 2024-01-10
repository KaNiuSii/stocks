package stocks

import (
	"math/rand"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) Rotate(deg float64) {
	rad := deg * math.Pi / 180
	x, y := v.X, v.Y
	v.X = x*math.Cos(rad) - y*math.Sin(rad)
	v.Y = x*math.Sin(rad) + y*math.Cos(rad)
}

func RandomVector2D() *Vector2D {
	return &Vector2D{X: 2 * rand.Float64() - 1, Y: 2 * rand.Float64() - 1}
}
