package main

import "fmt"

func main() {
	// variables are strict typed
	name := "Jeroen"

	fmt.Println(swap(name, "hello"))
}

// functions could return one or more arguments
func swap(x string, y string) (string, string) {
	return y, x
}
