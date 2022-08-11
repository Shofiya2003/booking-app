// in go everything has to be in a package. the go application has to be initialized first to create a mod file and every file must start with the package
package main

//Importing Format package from GO library
import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	//specifying types after the variable name
	var remainingTickets uint = 50
	//slice for all bookings
	var bookings []string
	var username string
	var userTickets uint

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have only %v tickets avalaible. Hurry!\n", remainingTickets)
	fmt.Printf("You can book tickets for %v here\n", conferenceName)

	for {

		//taking user input
		fmt.Printf("Enter your name\n")
		fmt.Scan(&username)

		fmt.Printf("Enter number of tickets you want to book\n")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, username)
		fmt.Printf("User %v booked %v tickets\n", username, userTickets)
		fmt.Printf("There are %v tickets left\n", remainingTickets)
		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("All are bookings are %v\n", firstNames)
	}

}
