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
	Status    string       // in progress / completed
	Type      string       // spending / income
	CancelFor *Transaction // cancelled transaction, nil - for normal transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []MCC) int64 {
	subsumsByMCC := make(map[MCC]int64)
	for _, transaction := range transactions {
		if transaction.CancelFor == nil {
			subsumsByMCC[transaction.MCC] += transaction.Sum
		} else {
			subsumsByMCC[transaction.CancelFor.MCC] -= transaction.Sum
		}
	}
	result := int64(0)
	for _, MCC := range mcc {
		result += subsumsByMCC[MCC]
	}
	return result
}
