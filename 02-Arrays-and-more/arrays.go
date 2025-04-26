package main

import "fmt"

func ArraysExample() {
	// Arrays in Go have a fixed length, because of that they're not used often.
	scores := [4]int{9001, 9333, 212, 33}

	for _, value := range scores {
		fmt.Println(value)
	}
	fmt.Println()
	scores[0] = 1009
	for _, value := range scores {
		fmt.Println(value)
	}

}
