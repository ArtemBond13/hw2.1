package card

import (
	"hw2.1/pkg/transfer"
)

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	Cards *[]Card
}

func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool){
	total = int64(0)
	isFrom := false
	isTo := false
	for _, cardFrom := range s.Cards{
		if cardFrom.Number == from{
			isFrom = true
		}
	}
	for _, cardTo := range s.Cards {
		if cardTo.Number == to {}
		isTo = true
	}
	if isFrom == true && isTo == true {
		transfer.NewService(&s.Cards, 0, 0)
	}
	return
}