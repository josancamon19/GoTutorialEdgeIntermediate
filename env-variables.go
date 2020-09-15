package main

import (
	"fmt"
	"os"
)

func main6() {
	fmt.Printf("Initial value: %s\n", os.Getenv("test-env"))
	os.Setenv("test-env", "Initial test env")
	fmt.Printf("Second value: %s\n", os.Getenv("test-env"))
	// Probably this was supposed to go in Beginner tutorials
}
