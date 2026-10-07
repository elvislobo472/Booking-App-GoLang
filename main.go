package main
import "fmt"  // fmt stands for format

func main() {

	var conferenceName = "LoelS3c Dev Conference"

	const conferenceTickets = 50
	var remainingTickets  = 50

	fmt.Print("Welcome to LoelS3c Ticketing System!!")
	fmt.Println("......Your place to get tickets for", conferenceName)

	fmt.Println()
	fmt.Println()

	fmt.Println("Total seats for the conference:", conferenceTickets) //without using format specifier

//Using format specifier

	fmt.Printf("There are %v seats remaining for the conference\n", remainingTickets)

	fmt.Print("Let's get you some seats :)")
}