package main

import (
	"fmt"
)

// var pow = []int{0, 2, 4, 6, 8, 10, 12, 14, 16}

func main() {
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// j := []struct {
	// 	i int
	// 	k bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(j)

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

	// var h []int
	// fmt.Println(h, len(h), cap(h))
	// if h == nil {
	// 	fmt.Println("nil!")
	// }

	// a := make([]int, 5)
	// printSlice("a", a)

	// b := make([]int, 0, 5)
	// printSlice("b", b)

	// c := b[:2]
	// printSlice("c", c)

	// d := c[2:5]
	// printSlice("d", d)

	// e := d[0:3]
	// printSlice("e", e)

	// // Create a tic-tac-toe board.
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// // The players take turns.
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"
	// board[2][0] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }
	// fmt.Println(len(board), cap(board))
	// var s []int
	// printSlice(s)

	// // append works on nil slices.
	// s = append(s, 5)
	// printSlice(s)

	// // The slice grows as needed.
	// s = append(s, 1)
	// printSlice(s)
	// s = append(s, 1)
	// printSlice(s)

	// // We can add more than one element at a time.
	// s = append(s, 2)
	// printSlice(s)

	// var f []int
	// printSlicef(f)

	// // append works on nil slices.
	// f = append(f, 5)
	// printSlicef(s)

	// // The slice grows as needed.
	// f = append(f, 1)
	// printSlicef(f)
	// f = append(f, 1)
	// printSlicef(f)

	// // We can add more than one element at a time.
	// f = append(f, 2)
	// printSlicef(f)

	// f = append(f, 2)
	// printSlicef(f)

	// f = append(f, 2)
	// printSlicef(f)

	// f = append(f, 2)
	// printSlicef(f)

	// f = append(f, 2)
	// printSlicef(f)

	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)
	// f = append(f, 2)
	// printSlicef(f)

	// for i, v := range pow {
	// 	fmt.Printf("2*%d = %d\n", i, v)
	// }

	powe := make([]int, 10)
	fmt.Println(powe)
	for i := range powe {
		powe[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(powe)
	for _, value := range powe {
		fmt.Printf("%d\n", value)
	}
	fmt.Println(powe)

}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// func printSlicef(f []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(f), cap(f), f)

// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

// }
