package main

import "fmt"

// functions could return one or more arguments
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	// GO has common types such as strings, ints and more
	// variables are strict typed
	name := "Jeroen"

	start, end := swap(name, "hello")
	fmt.Println(start, end)

	// Slices (Arrays) could be buffered and unbuffered
	list := [2]string{"Gopher", "Golang"}
	fmt.Println(list)

	// maps (also known as dictionaries) hold key/values
	dict := make(map[string]string)
	dict["world"] = "the earth, with all its countries, peoples and natural features"
	dict["computer"] = "an electronic machine that can store, organize and find information, do calculations and control other machines"

	fmt.Println(dict["world"])

	// you could loop over various types
	for key := range dict {
		fmt.Println(key)
	}

	// Loops could run endlessly
	for {
		fmt.Println("I am not a while loop.")
		break
	}
}
