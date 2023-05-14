package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}

/* To run this use
go run --filename
			or
go build --filename and then ./--filename
*/
