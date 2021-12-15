package main

import (
	"struct/management"
)

func main() {
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Fahmi"
	// user.LastName = "Alfareza"
	// user.Email = "fahmialfareza@icloud.com"
	// user.isActive = true

	// user := User{ID: 1, FirstName: "Fahmi", LastName: "Alfareza", Email: "fahmialfareza@icloud.com", isActive: true}

	// user := User{
	// 	ID: 1,
	// 	FirstName: "Fahmi", 
	// 	LastName: "Alfareza", 
	// 	Email: "fahmialfareza@icloud.com", 
	// 	IsActive: true,
	// }

	user1 := management.User{1, "Fahmi", "Alfareza", "fahmialfareza@icloud.com", true}
	user2 := management.User{2, "Fahmi", "Alfareza", "fahmialfareza@icloud.com", true}

	// displayUser1 := displayUser(user1)
	// displayUser2 := displayUser(user2)

	users := []management.User{user1, user2}

	group := management.Group{"Gamer", user1, users, true}

	// displayGroup(group)

	// result := user1.Display()

	// fmt.Println(result)

	group.Display()
}

// func displayUser(user User) string {
// 	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
// }

// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println()
// 	fmt.Printf("Member count: %d", len(group.Users))
// 	fmt.Println()

// 	fmt.Println("Users name: ")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }