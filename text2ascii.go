package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Passe uma string ao programa.")
		os.Exit(1)
	}

	for counter := 1; counter <= len(os.Args[1:]); counter++ {
		fmt.Printf("%s\t", os.Args[counter])
		for counter2 := 0; counter2 < len(os.Args[counter]); counter2++ {
			fmt.Printf("%d ", os.Args[counter][counter2])
		}
		fmt.Println()
	}
}
