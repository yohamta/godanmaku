package util

import "math"

// DegreeToDirectionIndex convert degree into 1 to 8 integer
func DegreeToDirectionIndex(degree int) int {
	adjust := 22.5
	return int(float64(degree)+90.0+360.0+adjust) % 360 / 45
}

// RadToDeg converts radian to degree
func RadToDeg(radian float64) int {
	return int(radian * 180 / math.Pi)
}

// DegToRad converts degree to radian
func DegToRad(degree int) float64 {
	return float64(degree) * math.Pi / 180
}
