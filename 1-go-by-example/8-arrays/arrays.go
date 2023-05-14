package main

import "fmt"

func main() {

	// Create an array with the length of 5
	var a [5]int
	fmt.Println("emp:", a) // [0,0,0,0]

	// setting 4th item in the array 100
	a[4] = 100
	fmt.Println("set:", a)    // [0 0 0 100]
	fmt.Println("get:", a[4]) // 100

	fmt.Println("len:", len(a)) // 5

	b := [5]int{1, 2, 3, 4, 5} // Short-hand declaration in a single line
	fmt.Println("dcl:", b)     // [1 2 3 4 5]

	var twoD [2][3]int // two dimensional array..
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[[0 1 2] [1 2 3]]]
}

/* To run this use
go run --filename
			or
go build --filename and then ./--filename
*/
