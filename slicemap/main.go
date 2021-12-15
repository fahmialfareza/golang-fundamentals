package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "E"},
	}

	// fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}