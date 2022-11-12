package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

    fmt.Printf("Welcome to our %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	bookings[0] = userFirstName + " " + userLastName

	remainingTickets -= uint(userTickets);

	fmt.Printf("Thanks %v %v for booking %v tickets. You will receive your confirmation at %v.\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
	fmt.Printf("%v",bookings[0])
}