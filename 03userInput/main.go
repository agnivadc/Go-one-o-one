package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate the pizza!!")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}