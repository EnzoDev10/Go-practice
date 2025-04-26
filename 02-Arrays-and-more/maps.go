package main

import "fmt"

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

// Son como los diccionarios de Python.
func Maps() {
	// Tambien se usa la función make para crearlos. El tamaño inicial(100) es opcional. Definir el tamaño puede mejorar la eficiencia.
	lookup := make(map[string]int, 100)
	lookup["goku"] = 9001

	// Tambien se pueden declarar de esta forma, llamada composite literal.
	lookupTwo := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	// Si se asigna el valor de una key a dos variables la primera contiene su valor y la segunda un booleano que confirma su existencia.
	power, exists := lookup["goku"]
	powerTwo, existsTwo := lookupTwo["gohan"]

	fmt.Println(power, exists)
	fmt.Println(powerTwo, existsTwo)

	// len retorna la cantidad total de keys.
	total := len(lookup)
	fmt.Println(total)

	// Delete elimina un valor de un map basado en su key.
	delete(lookup, "goku")

	// Como usar maps en structs.
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	fmt.Println(goku)
	fmt.Println()

	// se pueden iterar con un for + range. una variable para solo las keys y dos para los valores.
	for key := range lookupTwo {
		fmt.Println(key)
	}

}
