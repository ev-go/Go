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

func main() {
	// b := 1
	// y := 2
	// b, y = sw(b, y)
	// fmt.Println(b, y)

	b := 1
	y := 2
	sw(&b, &y)
	fmt.Println(b, y)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
