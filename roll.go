package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll() {
	fmt.Println("To get started, how many dice would you like to roll?")
	// Now that we're here from Main, let's initialize our number of dice.
	var numDice int
	// And the number of sides per dice. After all, this is a Polyhedral set we're playing with.
	var numSide int
	// Take in that data...
	fmt.Scan(&numDice)
	fmt.Println("Good, how many sides?")
	// That's some mighty good data
	fmt.Scan(&numSide)
	fmt.Printf("Alright, I'll roll %dd%d for you.\n", numDice, numSide)
	// Roll it! Roll them dice!
	for i := 1; i <= numDice; i++ {
		dieRoll := rand.Intn(numSide) + 1
		fmt.Println(dieRoll)
	}
	fmt.Println("Thank you for using the GoRoller.")
}

// Rolls a this just sets things up. The seed is initialized by the system clock.
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the GoRoller.")
	roll()
}
