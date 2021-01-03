package main

import (
	"fmt"
	"hw2.1/pkg/card"
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
		Number:   "5322 6729 6287 6282",
		Icon:     "http://",
	}
	visa2 := tinkoff.IssuerCard(0002, "Visa", 45679, "2345 4572 2845 8472")
	master2 := tinkoff.IssuerCard(0003, "Visa", 478312, "8765 3687 7462 8472")

	tinkoff.Cards = append(tinkoff.Cards, &visa)
	for _, val := range tinkoff.Cards {
		fmt.Println(val)
	}
	fmt.Println("")

	tinkoff.Cards = append(tinkoff.Cards, visa2, master2)
	for _, val := range tinkoff.Cards {
		fmt.Println(val)
	}

	//fmt.Println(visa, visa2, master2)
	//fmt.Println(tinkoff.Cards)

	fmt.Println(tinkoff.SearchByNumber("5322 6729 6287 6282"))
}
