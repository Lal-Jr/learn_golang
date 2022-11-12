package main

import "strings"

func validateInput(firstName string, lastName string, email string, tickets int, remainingTickets uint) (bool,bool,bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2;
	isValidEmail := strings.Contains(email, "@");
	isValidTicket := tickets > 0 && tickets <= int(remainingTickets);

	return isValidName, isValidEmail, isValidTicket
}