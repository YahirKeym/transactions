package main

import (
	"log"

	Mail "github.com/YahirKeym/transactions/internal/mail"
	GetTransaction "github.com/YahirKeym/transactions/internal/transaction/get"
	SaveTransaction "github.com/YahirKeym/transactions/internal/transaction/save"
	SummarizeTransaction "github.com/YahirKeym/transactions/internal/transaction/summarize"
	Database "github.com/YahirKeym/transactions/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db, err := Database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	transactions, err := GetTransaction.GetTransaction()
	if err != nil {
		log.Fatal(err)
	}
	err = SaveTransaction.Save(db, transactions)
	if err != nil {
		log.Fatal(err)
	}

	summaryTransaction := SummarizeTransaction.Summarize(transactions)
	err = Mail.Send(summaryTransaction, Mail.Email{To: "yakymaws@gmail.com"})
	if err != nil {
		log.Fatal(err)
	}
}
