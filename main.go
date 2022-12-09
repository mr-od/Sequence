package main

import (
	"fmt"
	"strconv"
	"strings"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i

	}
	return a
}

func addComma() string {
	a := makeRange(1, 13)
	var stringSlice []string
	for _, mr := range a {
		s := strconv.Itoa(mr)

		stringSlice = append(stringSlice, s)
	}

	var s = strings.Join(stringSlice, ".jpeg|")

	return s
}

func main() {
	s := addComma()

	jpg := strings.Split(s, "|")
	jpgs := strings.Join(jpg, ",")
	// jpgss := strings.Split(jpgs, ".")

	fmt.Println(jpgs)
}
