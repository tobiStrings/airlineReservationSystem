package airline

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateSeats(seat [10]bool) int{
	rand.Seed(time.Now().UnixNano())
	min:=0
	max:= 10
	seatNumber := rand.Intn(max-min)
	if seat[seatNumber] == false {
	}else {
		GenerateSeats(seat)
	}
	return seatNumber
}

func AssignFirstClassSeat(seatNumber, input int, seats [10]bool)string {
	if input == 1 && seatNumber <= 5 {
		seats[seatNumber] = true
		fmt.Println("You have been assigned a firstClass seat number",seatNumber+1)
		return "is generated"
	}else if seatNumber > 5{
		seatNumber = GenerateSeats(seats)
		AssignFirstClassSeat(seatNumber,input,seats)
		return "is generated"
	}
	return ""
}

func AssignEconomySeat(seatNumber, input int, seats [10]bool)string {
	if input == 2 && seatNumber > 5 {
		seats[seatNumber] = true
		fmt.Println("You have been assigned a ECONOMY seat number",seatNumber+1)
		return "is generated"
	}else if seatNumber <= 5{
		seatNumber = GenerateSeats(seats)
		AssignEconomySeat(seatNumber,input,seats)
		return "is generated"
	}
	return ""
}