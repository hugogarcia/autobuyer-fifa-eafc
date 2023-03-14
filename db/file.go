package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
)

var data = map[string]uint64{}

func SaveTrades() {
	logger.LogMessage(nil, "WriteTrades - Saving trades...")
	for tradeId, playerId := range entity.TradePlayerMap {
		data[fmt.Sprintf("%d", tradeId)] = playerId
	}

	if len(entity.TradePlayerMap) == 0 {
		return
	}

	j, err := json.Marshal(data)
	if err != nil {
		logger.PanicIt(err)
	}

	err = os.WriteFile("trades.json", j, 0644)
	if err != nil {
		logger.PanicIt(err)
	}
}

func ReadJsonFile() {
	content, err := os.ReadFile("trades.json")
	if err != nil {
		logger.PanicIt(err)
	}

	var trades map[string]uint64
	err = json.Unmarshal(content, &trades)
	if err != nil {
		logger.PanicIt(err)
	}

	if len(trades) == 0 {
		return
	}

	logger.LogMessage(nil, "Loading trades from file...")
	for tradeId, playerId := range trades {
		v, err := strconv.ParseUint(tradeId, 10, 64)
		if err != nil {
			logger.LogMessage(nil, "Error converting", err)
		}

		entity.TradePlayerMap[v] = playerId
	}
	logger.LogMessage(nil, "Trades loaded")
}
