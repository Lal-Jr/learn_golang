package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUser(conferenceName,conferenceTickets,remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets int
		
		fmt.Print("Enter First Name: ")
		fmt.Scan(&userFirstName)

		fmt.Print("Enter Last Name: ")
		fmt.Scan(&userLastName)

		fmt.Print("Enter Email Address: ")
		fmt.Scan(&userEmail)

		fmt.Print("Enter No. of Tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(userFirstName) > 2 && len(userLastName) > 2;
		isValidEmail := strings.Contains(userEmail, "@");
		isValidTicket := userTickets > 0 && userTickets <= int(remainingTickets);

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets -= uint(userTickets);
			bookings = append(bookings,userFirstName + " " + userLastName)

			firstName := getFirstNames(bookings)

			fmt.Printf("Thanks %v %v for booking %v tickets. You will receive your confirmation at %v.\n", userFirstName, userLastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
			fmt.Printf("they are all the bookings: %v.\n",firstName)

			if remainingTickets == 0 {
				fmt.Printf("%v is sold out",conferenceName)
				break
			}
			continue
		}
		fmt.Println("Invalid Name, Email or No. of Tickets")
	}
}

func greetUser(confName string, confTicket int, remTicket uint){
	fmt.Printf("Welcome to our %v booking application\n", confName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", confTicket, remTicket)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string{
	firstNames := []string{}
	for _,booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}