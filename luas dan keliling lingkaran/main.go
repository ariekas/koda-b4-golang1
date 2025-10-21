package main

import "fmt"

func circle(pi float64, r float64) (float64, float64) {
	area := pi * (r * r)
	circumference := 2 * pi * r
	return area, circumference
}

func main() {
	var pi, r float64

	fmt.Print("Masukkan nilai Pi: ")
	fmt.Scan(&pi)
	fmt.Print("Masukkan nilai jari-jari (r): ")
	fmt.Scan(&r)

	area, circumference := circle(pi, r)

	fmt.Println("Luas lingkaran:", area)
	fmt.Println("Keliling lingkaran:", circumference)
}
