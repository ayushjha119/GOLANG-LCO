package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my playground!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //notie b is small in bufio N is capital in NewReader o is capital in os S is capital in Stdin this is to which is public and which is private
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok syntx  or error  ok syntax

	input, _ := reader.ReadString('\n'); //something like try cath
	//input, err := reader.ReadString('\n');
	// _,err := reader.ReadString('\n'); if we care only about error	
	fmt.Println("Thanks for rating, ", input)

	fmt.Printf("Type of this rating is %T", input)





}