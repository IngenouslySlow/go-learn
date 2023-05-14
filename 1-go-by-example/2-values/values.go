package main

import "fmt"

func main() {
	// String concatenation
	fmt.Println("go" + "lang")

	// Number
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean
	fmt.Println("True & False = ", true && false)
	fmt.Println("True || False = ", true || false)
	fmt.Println("!true = ", !true)

}

/* To run this use
go run --filename
			or
go build --filename and then ./--filename
*/
