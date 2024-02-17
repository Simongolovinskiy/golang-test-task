package serializer

import (
	"encoding/json"
)

type CoinInfo struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float32 `json:"current_price"`
}

type Serializer struct{}

func (s Serializer) ConvertJson(data string) ([]CoinInfo, error) {
	var coinInfoList []CoinInfo

	err := json.Unmarshal([]byte(data), &coinInfoList)
	if err != nil {
		return nil, err
	}
	return coinInfoList, nil
}

func (s Serializer) GetCoinInfo(coinName string, data []CoinInfo) CoinInfo {
	for coin := 0; coin < len(data); coin++ {

		currCoin := data[coin]
		if currCoin.Name == coinName {
			return currCoin
		}

	}
	return CoinInfo{}
}
