package airline

import "testing"

func TestGenerateSeats(t *testing.T)  {
	var seats[10] bool
	seatNumber := GenerateSeats(seats)
	if seatNumber < 0 {
		t.Error("Expected seatNumber greater than or equals zero, but got",seatNumber)
	}
}

func TestAssignFirstClassSeat(t *testing.T) {
	var seats[10] bool
	str:=AssignFirstClassSeat(1,1,seats)
	if str != "is generated" {
		t.Error("Expected it is generated but got",str)
	}
}

func TestAssignFirstClassSeat2(t *testing.T) {
	var seats[10] bool
	str:=AssignFirstClassSeat(1,2,seats)
	if str != "" {
		t.Error("Expected it \"\" but got",str)
	}
}

func TestAssignEconomySeat(t *testing.T) {
	var seats[10] bool
	str:=AssignEconomySeat(7,2,seats)
	if str != "is generated" {
		t.Error("Expected it is generated but got",str)
	}
}

func TestAssignEconomySeat2(t *testing.T) {
	var seats[10] bool
	str:=AssignEconomySeat(8,1,seats)
	if str != "" {
		t.Error("Expected  \"\" but got",str)
	}
}