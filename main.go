package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y, z int) int {
	return x + y - z
}

func main() {

	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("My favorite %g number is. \n", math.Sqrt(9), rand.Intn(10))
	fmt.Printf("%g Welcome to the playground!", rand.Intn(100))
	fmt.Println("The time is", time.Now())

	fmt.Printf("Now you %g have problems.\n", math.Sqrt(9))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 15, 17))
}
