package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "eur"}
	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		log.Fatal(err)
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]
	fmt.Println(fmt.Sprintf("Bitcoin is worth %f usd (eur %f)", bitcoin["usd"], bitcoin["eur"]))
	fmt.Println(fmt.Sprintf("Ethereum is worth %f usd (eur %f)", eth["usd"], eth["eur"]))
}
