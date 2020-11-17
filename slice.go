package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	j := []struct {
		i int
		k bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(j)

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[0:]
	// fmt.Println(s)

	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// // Slice the slice to give it zero length.
	// s = s[:6]
	// printSlice(s)

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)

	var h []int
	fmt.Println(h, len(h), cap(h))
	if h == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	e := d[0:3]
	printSlice("e", e)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

// }
