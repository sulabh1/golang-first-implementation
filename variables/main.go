package main

import "fmt"

//declaring constant. capital letter indicated that its public

const Sulabh string = "Adhikari"

func main(){

	//string
	var username string = "sulabh"
	fmt.Println(username)
	fmt.Printf("This is of type %T \n",username)

	//integer without type
	var age = 24
	fmt.Println(age)
	fmt.Printf("This is of type %T \n", age)

	//integer with type uint unint 8 256 should be lss than 256
	var age1 uint8 = 255
	fmt.Println(age1)
	fmt.Printf("This is of type %T \n", age1)

	//bool
	var isLoggedIn bool = true
	if(isLoggedIn){
		fmt.Println(username, "Logged in true")
		fmt.Printf("This is of type %T\n",isLoggedIn)
	}else{
		fmt.Println(username, "Logged in false")
		fmt.Println(false)
	}

	//float same as int
	
	var price float64 = 29.16113434545433333
	fmt.Println(price, "price")
	fmt.Printf("This is of type %T \n", price)

	//value of variable after initilizing
	var anotherNumber int
	//only initilizing the value variable takes is 0
	fmt.Println(anotherNumber, "another number")
	fmt.Printf("This is of type %T\n",anotherNumber)

	//checking for string
	var anotherString string
	//empty string
	fmt.Println(anotherString, "another string")
	fmt.Printf("This is of type %T\n",anotherString)

	//implicit type means without declaring type. lexer automatically detect the type
	var website="my name is sulabh, and my website is swikritadhikari.com.np"
	fmt.Println(website, "implicit")

	//novar style:not allowed outside the function/method

	numberOfUsers :=399999
	fmt.Println(numberOfUsers)

	fmt.Println(Sulabh)


}


