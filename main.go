package main

import "fmt"

func main() {

	var conferenceName = "LoelS3c Dev Conference"

	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Print("********************Welcome to LoelS3c Ticketing System!!********************\n")

	fmt.Println()
	fmt.Println("Your place to get tickets for", conferenceName)

	fmt.Println()
	fmt.Println()

	fmt.Println("Total seats for the conference:", conferenceTickets) //without using format specifier

	//Using format specifier

	fmt.Printf("There are %v seats remaining for the conference\n", remainingTickets)
	fmt.Println()

	fmt.Print("Let's get you some seats :)\n")

	var userName string
	var userTickets, age uint
	var firstName, lastName, email string

	//Ask for username

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Printf("Dear %s, \n    You are not eligible to buy the tickets as your age is less than 18", firstName)
		return
	}

	fmt.Print("Enter your email ID: ")
	fmt.Scan(&email)

	fmt.Print("Enter your username: ")
	fmt.Scan(&userName)

	fmt.Print("How many tickets would you like to buy? ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email on %s\n\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v seats remaining for the conference!!\n", remainingTickets)

}
