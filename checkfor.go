package main

import (
	"fmt"
	"runtime"
	"time"
)

// func Sqrt(x, z float64) float64 {
// 	if v :=
// }
func switchtest(x, n int) string {
	v := "null"
	switch x {
	case n + 0:
		v = "equal"
	case n + 1:
		v = "less by 1"
	case n + 2:
		v = "less by 2"
	default:
		v = "Too far away."
	}
	return v
}

func main() {
	z := float64(100)
	x := float64(10)
	for i := 0; i < 7; i++ {
		z -= (z*z - x) / (2 * z)

		fmt.Println(z)
	}
	fmt.Println(z)
	fmt.Println(z * z)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// defer fmt.Println("world")

	// fmt.Println("hello")

	fmt.Println("counting")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	fmt.Println("done")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("switch tests")
	fmt.Println(switchtest(3, 2))
}
