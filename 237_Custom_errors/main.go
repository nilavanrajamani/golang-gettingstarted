package main

import "fmt"

type customErr struct {
	info string
}

func (c1 customErr) Error() string {
	return fmt.Sprintln("Error occured for the info", c1.info)
}

func foo(err error) {
	fmt.Println("Error is ", err)
}

func main() {
	custErr := customErr{info: "Sample error"}
	foo(custErr)
}
