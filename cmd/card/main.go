package main

import (
	"fmt"
	"github.com/artrey/bgo-std-card/pkg/card"
)

func main() {
	defaultCard := card.Card{
		Id:       1,
		Issuer:   "Visa",
		Balance:  25_453_00,
		Currency: "RUB",
		Number:   "4539709708333685",
		Icon:     "https://someIcon1.ico",
		Transactions: []*card.Transaction{
			{
				Id:        253,
				Sum:       1_203_91,
				Timestamp: 1610181644,
				MCC:       "5812",
				Status:    "completed",
			},
			{
				Id:        289,
				Sum:       2000_00,
				Timestamp: 1610181698,
				MCC:       "5912",
				Status:    "completed",
			},
			{
				Id:        411,
				Sum:       735_55,
				Timestamp: 1610181944,
				MCC:       "5411",
				Status:    "in progress",
			},
		},
	}
	fmt.Println(defaultCard)
	fmt.Println(defaultCard.Transactions[0])
	fmt.Println(card.SumByMCC(defaultCard.Transactions, []card.MCC{"5411", "5812"}))

	card.AddTransaction(&defaultCard, &card.Transaction{
		Id:        487,
		Sum:       1_105_00,
		Timestamp: 1610181990,
		MCC:       "5812",
		Status:    "in progress",
	})
	fmt.Println(card.SumByMCC(defaultCard.Transactions, []card.MCC{"5411", "5812"}))
}
