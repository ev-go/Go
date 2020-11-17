package main

import (
	"fmt"
)

// func Sqrt(x, z float64) float64 {
// 	if v :=
// }

// func sw(b, y int) (int, int) {
// 	return y, b
// }

func sw(b, y *int) {
	*b, *y = *y, *b
}

type vertex struct {
	X int
	Y int
}

var (
	v1 = vertex{1, 2}  // has type Vertex
	v2 = vertex{X: 1}  // Y:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	p  = &vertex{1, 2} // has type *Vertex
)

func main() {
	// b := 1
	// y := 2
	// b, y = sw(b, y)
	// fmt.Println(b, y)

	// b := 1
	// y := 2
	// sw(&b, &y)
	// fmt.Println(b, y)

	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j

	fmt.Println(vertex{1, 2})
	v := vertex{1, 2}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X, v.Y)
	// p := &v
	p.X = 25
	fmt.Println(v)

	fmt.Println(v1, *p, v2, v3)

	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}
