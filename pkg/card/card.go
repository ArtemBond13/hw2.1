package card

// Card абстракция банковской карты
type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

// Service хранятся карты банка
type Service struct {
	BankName string
	Cards    []*Card
}

func (s Service) IssuerCard(id int64, issuer string, balance int64, number string) *Card {
	card := &Card{
		Id:       id,
		Issuer:   issuer,
		Balance:  balance,
		Currency: "RUB",
		Number:   number,
		Icon:     "http://...",
	}
	s.Cards = append(s.Cards, card)
	return card
}

// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool) {
	total = 0
	//isFrom := false
	//isTo := false
	s.SearchByNumber(from)
	s.SearchByNumber(to)

	//if isFrom == true && isTo == true {
	//	transfer.NewService(s, 0, 0)
	//	if s
	//}
	return
}

// SearchByNumber поиска карты по номеру
func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}
