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

// NewService создает абстракцию банка
func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

// IssuerCard добавляет карту
func (s *Service) IssuerCard(id int64, issuer string, balance int64, number string) *Card {
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

// SearchByNumber поиска карты по номеру
func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}


func (s *Service) Add(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}