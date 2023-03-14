package search

import (
	"encoding/json"
	"fmt"

	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/request"
)

func GetTradeLiteById(tradesQuery string) entity.TradeResponse {
	path := "/ut/game/fifa23/trade/status/lite?tradeIds=" + tradesQuery
	body, err := request.GetBodyByPath(path)
	if err != nil {
		logger.PanicIt(err)
	}

	var trades entity.TradeResponse
	err = json.Unmarshal(body, &trades)
	if err != nil {
		logger.LogMessage(nil, body)
		logger.LogMessage(nil, err)
	}

	return trades
}

func GetTradeById(tradeId uint64) entity.Trade {
	trades := GetTradeLiteById(fmt.Sprintf("%d", tradeId))
	return trades.AuctionInfo[0]
}

func getNextBid(current uint32) uint64 {
	currentBid := uint64(current)
	nextBid := currentBid + 100
	if currentBid < 1000 {
		nextBid = currentBid + 50
	} else if currentBid >= 10000 {
		nextBid = currentBid + 250
	}

	return nextBid
}
