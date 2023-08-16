// favoriteFruit program asks users to input their favorite fruits and returns the phrases separated by a comma
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// function that asks the user a question and reads their input
func getUserInput(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + ": ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// gets rid of the new line from ReadString
	input = strings.TrimSpace(input)
	return input, nil
}

// function that asks if a user wishes to continue or end the program
func askUserToContinue() bool {
	// infinite for loop so user can input as many strings as they want as long as the conditional equals true.
	for {
		answer, err := getUserInput("Do you want to add another favorite fruit? [Yes or No]")
		if err != nil {
			log.Fatal(err)
		}
		// converts users answer input to a lower case string
		answer = strings.ToLower(answer)
		// any answer of "n" or "no" will break out of the loop, otherwise the loop will continue indefinitely
		// if the user inputs any other string the question will repeat itself
		if answer == "n" || answer == "no" {
			return false
		} else if answer == "y" || answer == "yes" || answer == "" {
			return true
		}
	}
}

// main function that asks the user a question and appends the slice with the users answer
// outputs the users inputs from the slice to a string, with each input separated by a comma
func main() {
	again := true
	var fruits []string
	// if user answers yes to continuing, for loop executes again asking for the user to input another fruit
	for bool(again) {
		input, err := getUserInput("Please enter your favorite fruit")
		if err != nil {
			log.Fatal(err)
		}
		// appends the slice with the user's input
		fruits = append(fruits, input)
		again = askUserToContinue()
	}
	fmt.Println("Your favorite fruits are: ", strings.Join(fruits, ", "))
}
