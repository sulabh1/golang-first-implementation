# golang-first-implementation

https://docs.google.com/document/d/1ONjgRWpyu_YWcg4tS50OvWapFyy0r7mQGZ8fIXGV-9o/edit?usp=sharing


learning golang from scratch.

=>Capital letter indicate that it is public

const Sulabh string = "Adhikari"


=>You cannot do following in outside code. Variable can be used without var keyword but only inside the function

func main(){
numberOfUsers :=399999
}


=>Can be treated as dynamically typed. Means without using any type.Also known as implicit type.

var website="my name is sulabh, and my website is swikritadhikari.com.np"
       fmt.Println(website, "implicit")

=> Empty variable declaration
string
var anotherString string
fmt.Println(anotherString, "another string")
       fmt.Printf("This is of type %T\n",anotherString)
Output is “” another string
This is of type string


Int
var anotherString string
fmt.Println(anotherNumber, "another number")
       fmt.Printf("This is of type %T\n",anotherNumber)
Output is 0 another number
This is of type int

=>reads input from os
	Bufio is a package which allow us to read the data and also os capture the standard input from our console.

reader:= bufio.NewReader(os.Stdin)
       fmt.Println("Enter the rating for my new looks")
=> comma ok || comma error syntax:
	. Its like try catch in other programming language.
	. Data coming stores in input and if error occur the stores in a error if we dont care about err then _ is used. Same goes to input as well. _ can be used in input if we dont care about incoming data. And \n is whenever user enter then our input exit.

input, _:= reader.ReadString('\n')


       fmt.Println("The value is:",input)





