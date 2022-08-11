// in go everything has to be in a package. the go application has to be initialized first to create a mod file and every file must start with the package
package main

//Importing Format package from GO library
import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const totalTickets = 50

// specifying types after the variable name
var remainingTickets uint = 50

// slice for all bookings
var bookings []string

func main() {
	greet()
	for {
		if remainingTickets == 0 {
			fmt.Printf("We are out of tickets. Sorry!")
			return
		}

		//taking user input
		firstname, lastname, email, userTickets := takeUserInput()

		isValid := validateInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValid {
			//print first names
			printFirstNames()
			bookTickets(firstname, lastname, userTickets)
		} else {
			continue
		}
	}

}

func printFirstNames() {
	var firstNames []string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("All are bookings are %v\n", firstNames)
}

func bookTickets(firstname string, lastname string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstname+" "+lastname)
	fmt.Printf("User %v booked %v tickets\n", bookings[len(bookings)-1], userTickets)
	fmt.Printf("There are %v tickets left\n", remainingTickets)
}

func greet() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have only %v tickets avalaible. Hurry!\n", remainingTickets)
	fmt.Printf("You can book tickets for %v here\n", conferenceName)
}

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

func takeUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var userTickets uint
	var email string
	fmt.Printf("Enter your first name\n")
	fmt.Scan(&firstname)

	fmt.Printf("Enter your last name\n")
	fmt.Scan(&lastname)

	fmt.Printf("Enter your email address\n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets you want to book\n")
	fmt.Scan(&userTickets)
	return firstname, lastname, email, userTickets
}
