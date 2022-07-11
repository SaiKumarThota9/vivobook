package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUser()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isVaildName, isVaildEmail, isVaildUserTicketsNumber := vaildateUserInput(firstName, lastName, email, userTickets)

		if isVaildName && isVaildEmail && isVaildUserTicketsNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {

			if !isVaildName {
				fmt.Println("First Name or LAst Name you Entered is too short ")
			}
			if !isVaildEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isVaildUserTicketsNumber {
				fmt.Println("Number of tickets you entered is invaild")
			}
		}
	}
}

func greetUser() {

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets you entered is invaild")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func vaildateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isVaildName := len(firstName) >= 2 && len(lastName) >= 2
	isVaildEmail := strings.Contains(email, "@")
	isVaildUserTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isVaildName, isVaildEmail, isVaildUserTicketsNumber

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter Number Of Tickets you need: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining of %v\n", remainingTickets, conferenceName)
}
