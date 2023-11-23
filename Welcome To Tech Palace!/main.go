package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	br := "***********************************************************************************************"
	return br[:numStarsPerLine] + "\n" + welcomeMsg + "\n" + br[:numStarsPerLine]
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.ReplaceAll(oldMsg, "\n", "")
	oldMsg = strings.Trim(oldMsg, " ")
	return oldMsg
	panic("Please implement the CleanupMessage() function")
}

func main() {
	oldMessage := "**************************\n*    BUY NOW, SAVE 10%   *\n**************************"
	fmt.Println(oldMessage)
	fmt.Println(CleanupMessage(oldMessage))
}
