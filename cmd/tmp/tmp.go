package main

import (
	"fmt"
	"log"

	"my/addToAnki/internal/infrastructure/clients/free_dictionary"
	free_dictionary2 "my/addToAnki/internal/usecases/anki/free_dictionary"
)

func main() {
	client := free_dictionary.New()

	uc := free_dictionary2.New(client)
	res, err := uc.GetData("ledger")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
