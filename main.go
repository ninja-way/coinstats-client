package main

import (
	"fmt"
	"github.com/ninja-way/coinstats-client/coinstats"
	"log"
	"time"
)

func main() {
	coinClient, err := coinstats.NewClient(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	coins, err := coinClient.GetCoins(10)
	if err != nil {
		log.Fatal(err)
	}

	for _, coin := range coins {
		fmt.Println(coin.Info())
	}
}
