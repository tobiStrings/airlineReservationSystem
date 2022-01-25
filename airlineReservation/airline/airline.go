package airline

import (
	"math/rand"
	"time"
)

func GenerateSeats(seat [10]bool) (int,[10]bool) {
	rand.Seed(time.Now().UnixNano())
	min:=0
	max:= 10
	seatNumber := rand.Intn(max-min)
	if seat[seatNumber] == false {
		seat[seatNumber] = true
	}else {
		GenerateSeats(seat)
	}
	return seatNumber,seat
}
