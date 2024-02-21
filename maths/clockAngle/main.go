package main

import (
	"fmt"
	"math"
)

func clockAngle(hours, minutes int) float64 {
	// Ensure hours and minutes are within valid ranges
	hours = hours % 12
	minutes = minutes % 60

	// Calculate the angles
	hourAngle := float64(30*hours) + float64(minutes)/2
	minuteAngle := float64(6 * minutes)

	// Calculate the absolute difference between the two angles
	angle := math.Abs(hourAngle - minuteAngle)

	// Ensure the angle is the acute angle (not reflex angle)
	if angle > 180 {
		angle = 360 - angle
	}

	return angle
}

func main() {
	// Test cases
	fmt.Println(clockAngle(30, 0))
	fmt.Println(clockAngle(3, 0))  // 90.0
	fmt.Println(clockAngle(6, 0))  // 180.0
	fmt.Println(clockAngle(12, 0)) // 0.0
	fmt.Println(clockAngle(9, 45)) // 22.5
}
