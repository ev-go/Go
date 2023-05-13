package slices_test

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	bracketRow := "[(())[]()[()]]"
	ok, err := bracketValidator(bracketRow)
	fmt.Println(ok, err)
}

func bracketValidator(bracketRow string) (bool, error) {
	validationSlice := make([]int32, 0, 100)
	rBrOp := "("
	sBrOp := "["
	rBrCl := ")"
	sBrCl := "]"
	pointer := 0
	verified := false
	var prevBr int32
	for _, value := range bracketRow {
		if (value == int32(rBrOp[0])) || (value == int32(sBrOp[0])) {
			validationSlice = append(validationSlice, value)
			pointer++
		} else if value == int32(rBrCl[0]) {
			prevBr = validationSlice[pointer-1]
			if prevBr == int32(rBrOp[0]) {
				pointer--
				validationSlice = validationSlice[:pointer]
			} else {
				return false, nil
			}
		} else if value == int32(sBrCl[0]) {
			prevBr = validationSlice[pointer-1]
			if prevBr == int32(sBrOp[0]) {
				pointer--
				validationSlice = validationSlice[:pointer]
			} else {
				return false, nil
			}
		}
	}
	verified = true
	return verified, nil
}
