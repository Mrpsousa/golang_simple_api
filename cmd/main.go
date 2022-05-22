package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrpsousa/adapter/repository"
	"github.com/mrpsousa/usecase/process_transaction"
)

func main() {

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    100,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err)
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}
