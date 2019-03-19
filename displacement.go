package main

import (
	"fmt"
	"math"
)

func input() (float64, float64, float64) {
	var acc, vel, disp float64
	fmt.Println("Enter the Acceleration ")
	fmt.Scan(&acc)
	fmt.Println("Enter the Initial Velocity ")
	fmt.Scan(&vel)
	fmt.Println("Enter the Initial Displacement")
	fmt.Scan(&disp)
	return acc, vel, disp

}
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	function := func(t float64) float64 {
		s := s0 + v0*t + ((a * math.Pow(t, 2)) / 2)
		return s
	}
	return function

}

func main() {
	a, v0, s0 := input()
	var t float64
	S := GenDisplaceFn(a, v0, s0)
	fmt.Println("Enter the Time ")
	fmt.Scan(&t)
	fmt.Println("Displacement at time ", t, "is :", S(t))

}
