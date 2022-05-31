package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func main() {
	st := "1-1"
	s, e := StringSum(st)
	fmt.Println(s, e)
}

func StringSum(input string) (output string, err error) {

	firstOperandString, secondOperandString := "", ""

	// clear from whitespaces:
	out := []rune(input)
	newOut := make([]rune, 0)
	for k := range out {
		if !unicode.IsSpace(rune(input[k])) {
			newOut = append(newOut, out[k])
		}
	}

	// check ig input is empty:
	if len(newOut) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	// check if first operand is negative
	if newOut[0] != '-' {
		newOut = append([]rune{'+'}, newOut...)
	}

	// preparing input string for operand number validation:

	tmpStr := strings.ReplaceAll(string(newOut), "+", ",")
	tmpStr = strings.ReplaceAll(string(tmpStr), "-", ",")

	// Now when we split input string shuld get three item slice
	// to be shure that the number of operands is correct
	if len(strings.Split(tmpStr, ",")) != 3 {
		return "", fmt.Errorf("incorrect number of operands: %w", errorNotTwoOperands)
	}

	// Now we are sure that number of operands is two.
	// We are preparing operands:
	for i := len(newOut) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(newOut[i])) {
			firstOperandString = string(newOut[:i])
			secondOperandString = string(newOut[i:])
			break
		}
	}

	// checking if operands number is greater than 2(skipping first rune because it is a sign):
	for j := len(secondOperandString) - 1; j >= 1; j-- {
		_, err := strconv.Atoi(string(secondOperandString[j]))
		if err != nil {
			return "", fmt.Errorf("error in second operand: %w", err)
		}
	}

	// checking if firs and second iperands are not empty strings
	if len(firstOperandString) == 0 || len(secondOperandString) == 0 {
		return "", fmt.Errorf("incorrect number of operands: %w", errorNotTwoOperands)
	}

	// checking if the first operand has letters by converting it to integer
	firstItem, err := strconv.ParseInt(string(firstOperandString), 10, 64)
	if err != nil {
		return "", fmt.Errorf("error in first operand: %w", err)
	}

	// checking if second operand has letters by converting it to integer
	secItem, err := strconv.ParseInt(string(secondOperandString), 10, 64)
	if err != nil {
		err = fmt.Errorf("error in second operand: %w", err)
		return "", err
	}

	return strconv.Itoa(int(firstItem) + int(secItem)), err

}
