package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isVaildName := len(firstName) >= 2 && len(lastName) >= 2
	isVaildEmail := strings.Contains(email, "@")
	isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isVaildName, isVaildEmail, isVaildTicketNumber
}
