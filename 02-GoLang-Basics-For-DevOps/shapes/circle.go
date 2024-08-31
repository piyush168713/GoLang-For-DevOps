package shapes

import "math"

func AreaOfCircle(radius float64) float64 {
	return math.Pi*radius*radius
}


func DiameterOfCircle(radius float64) float64 {
	return 2*math.Pi*radius
}