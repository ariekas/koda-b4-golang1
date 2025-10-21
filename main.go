package main

import "fmt"

func circle(pi float64, r float64) (float64) {
	area := pi * (r * r)
	return area
}

func main() {
	area := circle(22/7,10)
	fmt.Println(area)
}
