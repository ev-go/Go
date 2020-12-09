package main

import (
	"fmt"
)

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years);", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }
// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string

type ErrNegativeSqrt float64

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
