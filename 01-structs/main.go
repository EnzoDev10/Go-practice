package main

func main() {
	// the & is a pointer or address to a value.
	goku := &Saiyan{
		Person: &Person{"Goku"},
	}
	goku.Introduce()
}
