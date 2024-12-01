package main

import "fmt"

const loginToken string = "24qdadad23"

func main(){
	var user string = "Agniva"
	fmt.Printf("Type of variable is : %T \n", user)
	fmt.Println(user)

	var isLoggedIn bool = true
	fmt.Printf("Type of variable is : %T \n", isLoggedIn)
	fmt.Println(isLoggedIn)

	var smallVal uint8 = 234
	fmt.Printf("Type of variable is : %T \n", smallVal)
	fmt.Println(smallVal)

	var smallFloat float32 = 234.2423959493
	fmt.Printf("Type of variable is : %T \n", smallFloat)
	fmt.Println(smallFloat)

	//default values and aliases
	var testVar string
	fmt.Printf("Type of variable is : %T \n", testVar)
	fmt.Println(testVar)

	//implicit var declare
	var implicitVar = "Hi there"
	fmt.Printf("Type of variable is : %T \n", implicitVar)
	fmt.Println(implicitVar)

	//without var keyword
	noVar := 2943
	fmt.Printf("Type of variable is : %T \n", noVar)
	fmt.Println(noVar)

	fmt.Println(loginToken)
	fmt.Printf("Type of variable is : %T \n", loginToken)
}