package transfer

import (
	"github.com/ArtemBond13/hw2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		PercentTransfer   float64
		MinAmountTransfer int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		// TODO: Add test cases.
		cardSvc := card.S
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:           tt.fields.CardSvc,
				PercentTransfer:   tt.fields.PercentTransfer,
				MinAmountTransfer: tt.fields.MinAmountTransfer,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
