package main

import "fmt"

func main() {
	// var languages[5]string 
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	// languages := [5]string{"Go", "Ruby", "Javascript", "C#", "Pyhton"}

	// languages := [5]string{
	// 	"Go",
	//  	"Ruby",
	// 	"Javascript", 
	// 	"C#", 
	// 	"Pyhton",
	// }

	languages := [...]string{
		"Go",
	 	"Ruby",
		"Javascript", 
		"C#", 
		"Pyhton",
		"Kotlin",
	}
	
	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("index :", index, " language:", lang)
	}
}