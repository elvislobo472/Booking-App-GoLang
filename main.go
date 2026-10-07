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

	fmt.Println("Total seats for the conference:", conferenceTickets)
	fmt.Println("Remaining seats for the conference:", remainingTickets)

	fmt.Print("Let's get you some seats :)")
}