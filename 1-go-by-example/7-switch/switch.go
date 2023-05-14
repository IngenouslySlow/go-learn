package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, " as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("it's a weeekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(2)
	whatAmI(true)
	whatAmI("Hello")

	fmt.Println(time.Now())
}

/* To run this use
go run --filename
			or
go build --filename and then ./--filename
*/
