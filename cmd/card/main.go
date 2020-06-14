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
				Timestamp: 1592173658,
				MCC:       "5812",
				Status:    "completed",
				Type:      "spending",
			},
			{
				Id:        289,
				Sum:       2000_00,
				Timestamp: 1592174658,
				Status:    "completed",
				Type:      "income",
			},
			{
				Id:        411,
				Sum:       735_55,
				Timestamp: 1592183658,
				MCC:       "5411",
				Status:    "in progress",
				Type:      "spending",
			},
		},
	}
	fmt.Println(defaultCard)
	fmt.Println(defaultCard.Transactions[0])
	fmt.Println(card.SumByMCC(defaultCard.Transactions, []card.MCC{"5411", "5812"}))

	card.AddTransaction(&defaultCard, &card.Transaction{
		Id:        487,
		Sum:       1_203_91,
		Timestamp: 1592183958,
		Status:    "in progress",
		Type:      "income",
		CancelFor: defaultCard.Transactions[0],
	})
	fmt.Println(card.SumByMCC(defaultCard.Transactions, []card.MCC{"5411", "5812"}))

	fmt.Println(card.TranslateMCC(defaultCard.Transactions[0].MCC))
	fmt.Println(card.TranslateMCC("1154"))

	fmt.Println(defaultCard)
	fmt.Println(card.LastNTransactions(&defaultCard, 3))
	fmt.Println(card.LastNTransactions(&defaultCard, 10))
	fmt.Println(defaultCard)
}
