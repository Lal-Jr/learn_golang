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

    fmt.Printf("Welcome to our %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		remainingTickets -= uint(userTickets);
		bookings = append(bookings,userFirstName + " " + userLastName)

		firstNames := []string{}
		for _,booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Thanks %v %v for booking %v tickets. You will receive your confirmation at %v.\n", userFirstName, userLastName, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
		fmt.Printf("they are all the bookings: %v.\n",firstNames)

		if remainingTickets == 0 {
			fmt.Printf("%v is sold out",conferenceName)
			break
		}
	}
}