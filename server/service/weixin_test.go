package service

import (
	"fmt"
	"testing"
)

func TestGetWxcardTicket(t *testing.T) {
	ticket := GetWxcardTicket("wx71bcddd77824f86f")
	fmt.Println(ticket)
}
