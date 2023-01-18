package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:="welcome to go programming"
	fmt.Println(welcome)

	//reads ta input from os
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for my new looks")

	//comma ok || comma err syntax: replacement of try catch
	//data comming stores in input error stores in err but we are not interested in err so can use _
	input, _:= reader.ReadString('\n')

	fmt.Println("The value is:",input)
}