package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate the pizza!!")
	input, _ := reader.ReadString('\n')

	num, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("added 1 to your rating ", num+1)
	}
	// fmt.Println("Thanks for rating, ", input)
}