package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y, z int) int {
	return x + y - z
}

func swap(x, y, z string) (string, string, string) {
	return y, z, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	small = big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("My favorite %g number is. \n", math.Sqrt(9), rand.Intn(10))
	fmt.Printf("%g Welcome to the playground!", rand.Intn(100))
	fmt.Println("The time is", time.Now())

	fmt.Printf("Now you %g have problems.\n", math.Sqrt(9))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 15, 17))

	a, b, q := swap("hello", "world", "ho")
	fmt.Println(a, b, q)

	fmt.Println(split(17))

	// python := true

	// var i int
	fmt.Println(i, c, python, java)

	var i, j int = 2, 3
	k := 6
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T, Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T, Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

	// var z int
	var e float64
	var r bool
	var t string

	fmt.Printf("%v %v %v %q\n", z, e, r, t)

	x := 3.2
	y := 4.4

	// i := 42
	f := math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
	fmt.Printf("%T %T %T %T\n", x, y, f, z)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")
	fmt.Printf("%T\n", pi)
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))

}
