package main

import "fmt"

func main() {
	// afficher tout les combinaison de trois chiffre

	for i := 0; i < 10; i++ {
		for a := 0; a < 10; a++ {
			for x := 0; x < 10; x++ {
				if i < a && a < x {

					fmt.Printf("%v%v%v,", i, a, x)
				}

			}
		}
	}

}
