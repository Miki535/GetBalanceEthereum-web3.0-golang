package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addres := r.FormValue("Address")

		client, err := ethclient.Dial("https://mainnet.infura.io/v3/891cf9eba06143f2af8a4ebac812d40e")
		if err != nil {
			fmt.Println("Помилка при підключенні до Ethereum мережі:", err)
			return
		}
		
		address := common.HexToAddress(addres)

		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			fmt.Println("Помилка при отриманні балансу:", err)
			return
		}

		data := struct {
			Result *big.Int
		}{
			Result: balance,
		}
		tpl.Execute(w, data)
		return
	}
	tpl.Execute(w, nil)
}
