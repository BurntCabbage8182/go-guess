package main

import (
	"fmt"
	"math/rand"
)

func main() {

	secret := rand.Intn(1000) + 1

	var guess int

	fmt.Println("please enter a number from 1 to 1000:")

	for {

		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("You should enter a number not string!")

			var discard string

			fmt.Scan(&discard)

			continue

		}

		if secret == guess {
			fmt.Println("Good, you guessed it")

			break

		} else if secret > guess {
			fmt.Println("it is too low ")

		} else {
			fmt.Println("it is too high")
		}
	}
}
