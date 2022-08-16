// in go everything has to be in a package. the go application has to be initialized first to create a mod file and every file must start with the package
package main

//Importing Format package from GO library
import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceName = "Go Conference"

const totalTickets = 50

// specifying types after the variable name
var remainingTickets uint = 50

// slice for all bookings
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greet()

	if remainingTickets == 0 {
		fmt.Printf("We are out of tickets. Sorry!")
		return
	}

	//taking user input
	firstname, lastname, email, userTickets := takeUserInput()

	isValid := validateInput(firstname, lastname, email, userTickets, remainingTickets)

	if isValid {
		bookTickets(firstname, lastname, email, userTickets)
		wg.Add(1)
		go printTicket(userTickets, firstname, lastname, email)
	} else {
		// continue
	}

	wg.Wait()

}

func printFirstNames() {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("All are bookings are %v\n", firstNames)
}

func bookTickets(firstname string, lastname string, email string, userTickets uint) {
	//create an empty map
	user := UserData{firstName: firstname, lastName: lastname, email: email, userTickets: userTickets}
	remainingTickets -= userTickets
	bookings = append(bookings, user)
	fmt.Printf("User %v booked %v tickets\n", firstname, userTickets)
	fmt.Printf("There are %v tickets left\n", remainingTickets)
}

func greet() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have only %v tickets avalaible. Hurry!\n", remainingTickets)
	fmt.Printf("You can book tickets for %v here\n", conferenceName)
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

func printTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets booked by %v %v", userTickets, firstname, lastname)
	fmt.Println("############################")
	fmt.Printf("sending email:\n %v \n to email address %v\n", ticket, email)
	fmt.Printf("Hope to see you there!")
	fmt.Println("############################")
	wg.Done()
}
