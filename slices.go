package main

import "fmt"

func main() {
	s := []string{}
	printSlice(s)

	// append works on nil slices.
	s = append(s, "one one")
	printSlice(s)

	// The slice grows as needed.
	s = append(s, "two")
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, "three", "four")
	printSlice(s)

	s = []string{}
	printSlice(s)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

