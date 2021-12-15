package main

import "fmt"

type Student struct {
	ID 		int
	Name 	string
	GPA 	float32	
}

func (student *Student) graduate() {
	student.Name = student.Name + ", S.Kom"
}

func main() {
	student := Student{1, "Fahmi Alfareza", 3.55}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}