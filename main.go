package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

    fmt.Printf("Welcome to our %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	
	fmt.Scan(&userName)
	userTickets = 2

	fmt.Printf("%v booked %v tickets", userName, userTickets)
}