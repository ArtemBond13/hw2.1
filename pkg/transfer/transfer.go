package transfer

import (
	"fmt"
	"github.com/ArtemBond13/hw2.1/pkg/card"
)

type Service struct {
	CardSvc           *card.Service
	PercentTransfer   float64
	MinAmountTransfer int64
}

func NewService(cardSVC *card.Service, percent float64, minAmount int64) *Service {
	return &Service{
		CardSvc:           cardSVC,
		PercentTransfer:   percent,
		MinAmountTransfer: minAmount,
	}
}

// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	total = 0
	ofFrom := s.CardSvc.SearchByNumber(from)
	onTo := s.CardSvc.SearchByNumber(to)
	cardService := NewService(s.CardSvc, 0.5, 10_00)

	commission := int64(float64(amount) * cardService.PercentTransfer)
	if commission < cardService.MinAmountTransfer {
		commission = cardService.MinAmountTransfer
	}

	if ofFrom == nil && onTo == nil {
		total = amount + commission
		ok = true
	} else if ofFrom == nil && onTo != nil {
		onTo.Balance += amount
		total = amount + commission
		ok = true
		fmt.Print(onTo.Balance, "\n")
	} else if ofFrom != nil && onTo == nil {
		total = amount + commission
		if ofFrom.Balance < total {
			ok = false
			return total, ok
		}
		ofFrom.Balance -= total
		ok = true
	} else {
		total = amount + commission
		if ofFrom.Balance < total {
			ok = false
			return total, ok
		}
		ofFrom.Balance -= total
		onTo.Balance += amount
		ok = true
	}
	return
}
