package SaveTransaction

import (
	"database/sql"

	GetTransaction "github.com/YahirKeym/transactions/src/internal/transaction/get"

	"github.com/google/uuid"
)

func Save(db *sql.DB, transactions []GetTransaction.Transaction) error {
	account, err := saveAccount(db)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO transactions (date, amount, accountId) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	for _, transaction := range transactions {
		_, err := stmt.Exec(transaction.Date, transaction.Amount, account)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveAccount(db *sql.DB) (string, error) {
	accountId := uuid.New().String()
	stmt, err := db.Prepare("INSERT INTO account (id) VALUES (?)")
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(accountId)
	if err != nil {
		return "", err
	}
	return accountId, nil
}
