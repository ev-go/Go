package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// func closesqrt(x float) float64 {
// 	if z
// }

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// fmt.Println(
	// 	closesqrt (41)
	// )

	z := 0
	x := 0
	for i := 0; i < 10; i++ {
		z++
		z -= (z*z - x) / (2 * z)
		fmt.Println(sum)
	}
	fmt.Println(sum)

}
