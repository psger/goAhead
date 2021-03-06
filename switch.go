package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // 可以多个条件
		fmt.Println("It`s the weekday")
	default:
		fmt.Println("It`s the workday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Println("Do not know type", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}
