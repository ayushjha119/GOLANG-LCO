package main

import "fmt"

// jwtToken := 300000 //error	//outside function we can't use this style

const LoginToken string = "efvdjvjkJK" //capital lelter means public // You can use anywhere

func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn  bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)


	var smallVal  uint8 = 255
	// var smallVal  uint8 = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat  float64 = 255.45621447852336545
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	//default values and some aliases

	var anotherVariable int; //deafault value is 0
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	var anottherstring string; //deafault value is empty string	
	fmt.Println(anottherstring)
	fmt.Printf("Variable is of type: %T\n", anottherstring)

	//implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	var numb = 123
	fmt.Println(numb)
	fmt.Printf("Variable is of type: %T\n", numb)

	//no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T\n", numberOfUser)

	strinOfuser := "HEllo"
	fmt.Println(strinOfuser)
	fmt.Printf("Variable is of type: %T\n", strinOfuser)



	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)



}