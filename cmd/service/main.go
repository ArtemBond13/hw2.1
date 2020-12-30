package main

import (
	"fmt"
	"hw2.1/pkg/card"
	"hw2.1/pkg/transfer"
)

func main() {

	cardSVC := transfer.NewService(&card.Service{})
	fmt.Println(cardSVC)
}
