package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.1/pkg/card"
)

func main() {

	//tinkoff := card.Service{
	//	BankName: "Tinkoff",
	//}

	tot := card.NewService("Tot")

	tot.Add(&card.Card{Balance: 34566, Number: "2344"}, &card.Card{Balance: 10000_00, Number: "2345"})
	fmt.Print(tot.SearchByNumber("2344"))
	//master := card.Card{
	//	Id:       0001,
	//	Issuer:   "MaaterCard",
	//	Balance:  45234,
	//	Currency: "",
	//	Number:   "3214 9324 2457 7365",
	//	Icon:     "http://",
	//}
	//
	//visa := card.Card{
	//	Id:       0001,
	//	Issuer:   "VISA",
	//	Balance:  2342234,
	//	Currency: "",
	//	Number:   "4539526713450817073",
	//	Icon:     "http://",
	//}
	//visa2 := tinkoff.IssuerCard(0002, "Visa", 45679, "4024007196899149")
	//master2 := tinkoff.IssuerCard(0003, "MasterCard", 478312, "5149611582430745")
}
