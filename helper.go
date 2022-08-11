package main

import (
	"fmt"
	"strings"
)

func validateInput(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) bool {
	var isValidName = len(firstname) > 2 && len(lastname) > 2
	var isValidUserTickets = userTickets > 0 && userTickets <= remainingTickets
	var isValidEmail = strings.Contains(email, "@")
	if !isValidEmail {
		fmt.Printf("Email is invalid\n")
	}
	if !isValidName {
		fmt.Printf("Name is invalid\n")
	}
	if !isValidUserTickets {
		fmt.Printf("Number of tickets is invalid\n")
	}
	return isValidEmail && isValidUserTickets && isValidName
}
