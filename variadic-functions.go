package main

import "fmt"

func variadicFunction(stringVars ...string) {
	fmt.Println(stringVars)
}

func variadicFunction2(anyVars ...interface{}) {
	fmt.Println(anyVars)
}

func main5() {
	variadicFunction("1", "2", "3")
	variadicFunction2("1", "2", "3", 1, 2, 3)
	// Probably this was supposed to go in Beginner tutorials

}
