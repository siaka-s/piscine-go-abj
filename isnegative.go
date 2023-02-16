package main

import "fmt"

func isnegative(nb int) {

	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}

}

func main() {
	// exécution de fonction avec affectation de valeur
	isnegative(-1)
}
