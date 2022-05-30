package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

// func main() {
// 	st := "  13  -4 "
// 	s, e := StringSum(st)
// 	fmt.Println(s, e)
// }

func StringSum(input string) (output string, err error) {

	firstAddenString, secondAddenString := "", ""

	// clear from whitespaces:
	out := []rune(input)
	newOut := make([]rune, 0)
	for k, _ := range out {
		if !unicode.IsSpace(rune(input[k])) {
			newOut = append(newOut, out[k])
		}
	}
	if len(newOut) == 0 {
		return "", errorEmptyInput
	}
	for i := len(newOut) - 1; i > 0; i-- {
		if !unicode.IsDigit(rune(newOut[i])) {
			if !unicode.IsDigit(rune(newOut[i-1])) {
				err = fmt.Errorf("wrong operation please use either + or -: %v", errorNotTwoOperands)
				return "", err
			}

			firstAddenString = string(newOut[:i])
			secondAddenString = string(newOut[i:])
			break
		}
	}

	firstItem, err := strconv.Atoi(string(firstAddenString))
	if err != nil {
		err = fmt.Errorf("%v, %v", err, errorNotTwoOperands)
		return "", err
	}

	secItem, err := strconv.Atoi(string(secondAddenString))
	if err != nil {
		err = fmt.Errorf("%v, %v", err, errorNotTwoOperands)
		return "", err
	}

	return strconv.Itoa(firstItem + secItem), err

}
