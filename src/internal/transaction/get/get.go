package GetTransaction

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	ConstantsTransactions "stori/src/internal/transaction"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Date   time.Time
	Amount float64
}

func GetTransaction() ([]Transaction, error) {
	filePath := ConstantsTransactions.Constants().SamplePathFile
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	transactions := []Transaction{}
	for {
		record, err := reader.Read()
		if len(record) == 0 {
			break
		}
		dateOffParse, transaction := record[1], record[2]
		if record[2] != "Transaction" {
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			amount, err := strconv.ParseFloat(strings.TrimPrefix(transaction, "+"), 64)
			if err != nil {
				return nil, err
			}
			dateWithYear := time.Now().Format("2006") + "/" + dateOffParse
			date, err := time.Parse("2006/1/2", dateWithYear)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, Transaction{Date: date, Amount: amount})
		}
	}
	return transactions, nil
}
