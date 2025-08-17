package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main (){
	// Generates the dice throws
	rand.Seed(time.Now().UnixNano())

	balance := 1000;
	for balance > 0 {
		//rolls the dice
		dice  := rand.Intn(6) + 1;

		var bet int
		fmt.Print("Enter the value that you want to bet")
		fmt.Scan(&bet)
		if bet > balance{
			fmt.Print("You don't have this amount of money!!!")
			continue
		}

		var guess int
		fmt.Print("Guess the dice roll (1-6)")
		fmt.Scan(&guess)

		if guess == dice {
			fmt.Println("🎉 Correct! You guessed it!")
		balance += bet*3 // reward
		} else {
			fmt.Printf("❌ Wrong! The dice was %d\n", dice)
		}
	}
	fmt.Printf("Game Over")
}