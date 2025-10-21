package main

import "fmt"

func circle(pi float64, r float64) (float64, float64) {
	area := pi * (r * r)
	circumference := 2*pi*r
	return area, circumference
}

func main() {
	area,circumference := circle(3.14,10)
	fmt.Println(area)
	fmt.Println(circumference)

}
