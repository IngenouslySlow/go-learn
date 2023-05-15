package main

import "fmt"

func main() {
	word := "Hello"

	for _, letter := range word {
		fmt.Println(letter)
	}
}

/* To run this use
go run --filename
			or
go build --filename and then ./--filename
*/
