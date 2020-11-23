package main

import (
	"fmt"
)

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bel Labs": {40.68433, -74.39967},
// 	"Google":   {37.42202, -122.08408},
// }

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

func fibonacci() func() int {
	f2, f1 := 0, 1
	return func() int {
		f := f2
		f2, f1 = f1, f+f1

		return f
	}
}

func main() {
	// m := make(map[string]int)

	// m["Answer"] = 42
	// fmt.Println("The value:", m["Answer"])

	// m["Answer"] = 48
	// fmt.Println("The value:", m["Answer"])

	// delete(m, "Answer")
	// fmt.Println("The value:", m["Answer"])

	// m["Answer"] = 0
	// delete(m, "Answer")
	// v, ok := m["Answer"]
	// fmt.Println("The value:", v, "Present?", ok)
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(2, 1))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
