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

		seatGenerated,seats := airline.GenerateSeats(seats)
		if input == 1 && seatGenerated <= 5{
			fmt.Println("You have been assigned a firstClass seat number",seatGenerated+1)
		}else if input ==2 && seatGenerated >5 {
			for false {
				seatGenerated,seats = airline.GenerateSeats(seats)
			}
			fmt.Println("You have been assigned an Economy seat number",seatGenerated+1)
		}


}
