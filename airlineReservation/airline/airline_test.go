package airline

import "testing"

func TestGenerateSeats(t *testing.T)  {
	var seats[10] bool
	seatNumber ,_:= GenerateSeats(seats)
	if seatNumber < 0 {
		t.Error("Expected seatNumber greater than or equals zero, but got",seatNumber)
	}
}

