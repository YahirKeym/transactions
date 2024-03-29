package SummarizeTransaction

import (
	"sort"
	"time"

	GetTransaction "github.com/YahirKeym/transactions/internal/transaction/get"
)

type SummaryTransaction struct {
	Balance             float64
	TransactionsByMonth map[time.Month]int
	AvgCredit           float64
	AvgDebit            float64
	SortMonths          []time.Month
	CreditTransactions  int
	DebitTransactions   int
}

func Summarize(transactions []GetTransaction.Transaction) SummaryTransaction {
	balance := 0.0
	transactionsByMonth := make(map[time.Month]int)
	totalCredit, totalDebit := 0.0, 0.0
	positiveTransactions := 0
	debitTransactions := 0

	for _, transaction := range transactions {
		balance += transaction.Amount
		month := transaction.Date.Month()
		transactionsByMonth[month] = transactionsByMonth[month] + 1
		if transaction.Amount > 0 {
			totalCredit += transaction.Amount
			positiveTransactions++
		} else {
			totalDebit += transaction.Amount
			debitTransactions++
		}
	}
	sortMonths := sortTransactionsByMonth(transactionsByMonth)
	avgCredit := average(totalCredit, positiveTransactions)
	avgDebit := average(totalDebit, len(transactions))
	summaryTransaction := SummaryTransaction{
		Balance:             balance,
		TransactionsByMonth: transactionsByMonth,
		AvgCredit:           avgCredit,
		AvgDebit:            avgDebit,
		SortMonths:          sortMonths,
		CreditTransactions:  positiveTransactions,
		DebitTransactions:   debitTransactions,
	}
	return summaryTransaction
}

func average(total float64, count int) float64 {
	if count == 0 {
		return 0
	}
	return total / float64(count)
}

func sortTransactionsByMonth(transactions map[time.Month]int) []time.Month {
	months := make([]time.Month, 0, len(transactions))
	for month := range transactions {
		months = append(months, month)
	}

	sort.Slice(months, func(i, j int) bool {
		return months[i] < months[j]
	})

	return months
}
