package main

import (
	"fmt"
	"golang-test-task/first_task/api"
	"golang-test-task/first_task/interfaces"
	"golang-test-task/first_task/serializer"
	"log"
)

func run() {
	var coinApi interfaces.ApiInt = api.CoinApi{}
	var coinSerializer interfaces.SerializerInt = serializer.Serializer{}
	coinList, err := coinSerializer.ConvertJson(coinApi.GetCoins())
	if err != nil {
		log.Fatal("couldn't get coinList", err)
	}
	currentCoin := coinSerializer.GetCoinInfo("Ethereum", coinList)
	fmt.Println(currentCoin)
}

func main() {
	run()
}
