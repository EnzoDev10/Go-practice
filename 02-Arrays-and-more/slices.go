package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// los slices permiten usar funcionalidades de arrays sin una capacidad limite.
func Slices() {
	/* 	// Make crea un slice declarando un tipo,la cantidad de items que existen dentro y su capacidad maxima.
	   	scores := make([]int, 0, 5)
	   	// cap muestra la capacidad maxima de un array o slice.
	   	c := cap(scores)
	   	fmt.Println(c)
	   	for i := 0; i < 25; i++ {
	   		// Append puede aumentar la capacidad. Los slices no aumentan su capacidad de uno en uno si no que duplican la capacidad actual.
	   		// ej. capacidad de 3 y len de 3 luego de un append la capacidad pasa a ser de 6 y el len de 4.
	   		scores = append(scores, i)

	   		if cap(scores) != c {
	   			c = cap(scores)
	   			fmt.Println(c)
	   		}
	   	} */

	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:3])
	fmt.Println(scores[:3])
	fmt.Println(worst)
}
