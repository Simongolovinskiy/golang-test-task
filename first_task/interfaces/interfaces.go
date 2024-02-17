package interfaces

import "golang-test-task/first_task/serializer"

type ApiInt interface {
	GetCoins() string
}

type SerializerInt interface {
	ConvertJson(data string) ([]serializer.CoinInfo, error)
	GetCoinInfo(coinName string, data []serializer.CoinInfo) serializer.CoinInfo
}
