package main

import (
	"airlineReservationSystem/airlineReservation/airline"
	"fmt"
)

func main() {
	var seats [10]bool
	var input int

		fmt.Println("Enter 1 for a first class seat or 2 for Economy seat")
		fmt.Scanln(&input)

		seatGenerated := airline.GenerateSeats(seats)
	if input == 1 {
		airline.AssignFirstClassSeat(seatGenerated,input,seats)
	}else if input == 2{
		airline.AssignEconomySeat(seatGenerated,input,seats)
	}
}
