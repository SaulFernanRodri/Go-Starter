package main

import "fmt"

func main() {
	// If-else statement
	num := 10
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	// Declare the map
	someMap := map[string]int{
		"key": 42,
		"foo": 100,
		"bar": 200,
	}

	// Comma ok idiom
	if value, ok := someMap["key"]; ok {
		fmt.Println("Key found:", value)
	} else {
		fmt.Println("Key not found")
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("Midweek day")
	}

	// Select statement
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "message from ch1"
	}()

	go func() {
		ch2 <- "message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received", msg2)
	}
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop as a while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop with break
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	// Continue statement
	for l := 0; l < 5; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}
}
