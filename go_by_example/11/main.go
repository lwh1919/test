package main

import "errors"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var ErrDivideByZero = errors.New("divide by zero")

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {}
