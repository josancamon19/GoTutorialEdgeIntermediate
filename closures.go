package main

import "fmt"

func getLimit() func() int {
	// a closure is a function value which is able to reference variables that lay outwith it’s body.
	// https://tutorialedge.net/golang/go-closures-tutorial/
	limit := 10
	return func() int {
		// In this case, this function returned is using an outside variable
		limit -= 1
		return limit
	}
}

func main() {
	limitFunction := getLimit()
	fmt.Println(limitFunction())
	fmt.Println(limitFunction())
	fmt.Println(limitFunction())
	limitFunction2 := getLimit()
	fmt.Println(limitFunction2())

}
