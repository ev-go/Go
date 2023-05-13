package slices_test

import (
	"fmt"
	"testing"
)

type Validator interface {
	Put(element string)
	Get() string
	Pop()
}

func NewValidator() Validator {
	return &Stack{
		validationSlice: make([]string, 0, 10),
	}
}

type Stack struct {
	validationSlice []string
}

func (stack *Stack) Put(element string) {
	stack.validationSlice = append(stack.validationSlice, element)
}

func (stack *Stack) Get() string {
	return stack.validationSlice[len(stack.validationSlice)-1]
}

func (stack *Stack) Pop() {
	stack.validationSlice = stack.validationSlice[:len(stack.validationSlice)-1]
}

func bracketValidatorPkg(inputRow string) bool {
	validator := NewValidator()
	rBrOp := "("
	sBrOp := "["
	rBrCl := ")"
	sBrCl := "]"

	var valueStr, prevBracket string
	for _, value := range inputRow {
		valueStr = string(value)
		if (valueStr == rBrOp) || (valueStr == sBrOp) {
			validator.Put(valueStr)
		} else if valueStr == rBrCl {
			prevBracket = validator.Get()
			if prevBracket == rBrOp {
				validator.Pop()
			} else {
				return false
			}
		} else if valueStr == sBrCl {
			prevBracket = validator.Get()
			if prevBracket == sBrOp {
				validator.Pop()
			} else {
				return false
			}
		}
	}
	return true
}

func TestSlices(t *testing.T) {
	bracketRow := "[(())[]()[()]]"
	ok := bracketValidatorPkg(bracketRow)
	fmt.Println(ok)
}

func bracketValidator(bracketRow string) bool {
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
				return false
			}
		} else if value == int32(sBrCl[0]) {
			prevBr = validationSlice[pointer-1]
			if prevBr == int32(sBrOp[0]) {
				pointer--
				validationSlice = validationSlice[:pointer]
			} else {
				return false
			}
		}
	}
	verified = true
	return verified
}
