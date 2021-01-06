package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.1/pkg/card"
	"github.com/ArtemBond13/hw2.1/pkg/transfer"
)

func main() {

	tinkoff := card.Service{
		BankName: "Tinkoff",
	}
	//master := card.Card{
	//	Id:       0001,
	//	Issuer:   "MaaterCard",
	//	Balance:  45234,
	//	Currency: "",
	//	Number:   "3214 9324 2457 7365",
	//	Icon:     "http://",
	//}
	//
	visa := card.Card{
		Id:       0001,
		Issuer:   "VISA",
		Balance:  2342234,
		Currency: "",
		Number:   "4539526713450817073",
		Icon:     "http://",
	}
	visa2 := tinkoff.IssuerCard(0002, "Visa", 45679, "4024007196899149")
	master2 := tinkoff.IssuerCard(0003, "MasterCard", 478312, "5149611582430745")

	tinkoff.Cards = append(tinkoff.Cards, &visa)
	for _, val := range tinkoff.Cards {
		fmt.Println(val)
	}
	fmt.Println("")

	//tinkoff.Cards = append(tinkoff.Cards, visa2, master2)
	for _, val := range tinkoff.Cards {
		fmt.Println(val)
	}

	fmt.Println(visa, visa2, master2)
	//fmt.Println(tinkoff.Cards)

	//fmt.Println(tinkoff.SearchByNumber("5322 6729 6287 6282"))
	tinkof := transfer.Service{&tinkoff, 0, 10}
	fmt.Println(tinkof.Card2Card("453952671345081707", "4024007196899149", 100_000_00))
	// fmt.Println(tinkof.CardSvc.Cards)
}
