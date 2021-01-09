package card

type MCC string

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []*Transaction
}

type Transaction struct {
	Id        int64
	Sum       int64
	Timestamp int64
	MCC       MCC
	Status    string // in progress / completed / cancelled
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []MCC) int64 {
	subsumsByMCC := make(map[MCC]int64)
	for _, transaction := range transactions {
		subsumsByMCC[transaction.MCC] += transaction.Sum
	}
	result := int64(0)
	for _, MCC := range mcc {
		result += subsumsByMCC[MCC]
	}
	return result
}

func LastNTransactions(card *Card, n int) []*Transaction {
	transactionsCount := len(card.Transactions)

	if n > transactionsCount {
		n = transactionsCount
	}

	result := make([]*Transaction, 0, n)
	for i := transactionsCount - 1; i >= transactionsCount - n; i-- {
		result = append(result, card.Transactions[i])
	}
	return result
}
