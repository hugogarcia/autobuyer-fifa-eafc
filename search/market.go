package search

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/bid"
	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/request"
)

func SearchPlayers() {
	logger.LogMessage(nil, "searching players...")

	for playerId, v := range entity.Players {
		if v.MaxBid == 0 {
			fmt.Println("Player:", v.Name, "without bid")
			continue
		}

		searchByPlayer(playerId, v)
		time.Sleep(time.Second * 7)
	}

	SearchPlayers()
}

func searchByPlayer(id uint64, player entity.Player) {
	logger.LogMessage(nil, "Searching player:", player.Name)

	path := fmt.Sprintf("/ut/game/fc24/transfermarket?num=21&start=0&type=player&maskedDefId=%d&macr=%d", id, player.MaxBid)
	body, err := request.GetBodyByPath(path)
	if err != nil {
		logger.PanicIt(err)
	}

	var trades entity.TradeResponse
	err = json.Unmarshal(body, &trades)
	if err != nil {
		logger.LogMessage(nil, string(body[:]))
		logger.LogMessage(nil, err)
	}

	tradesToBid := filterAvailableTrades(trades, player.MaxBid)

	for _, v := range tradesToBid.AuctionInfo {
		nextBid := getNextBid(v.CurrentBid)
		if nextBid > player.MaxBid {
			continue
		}

		t := GetTradeById(v.TradeId)
		nextBid = getNextBid(t.CurrentBid)
		if t.CurrentBid == 0 {
			nextBid = getNextBid(t.StartingBid)
		}

		if nextBid > player.MaxBid {
			continue
		}

		bid.MakeBid(v.TradeId, nextBid, player.Name)
		time.Sleep(time.Second)
	}
}

func filterAvailableTrades(trades entity.TradeResponse, maxBid uint64) entity.TradeResponse {
	tradesFilter := entity.TradeResponse{
		Credits: trades.Credits,
	}

	for _, tr := range trades.AuctionInfo {
		if tr.BidState == "highest" ||
			tr.TradeState != "active" ||
			tr.Expires <= 0 ||
			tr.Watched ||
			tr.CurrentBid >= uint32(maxBid) ||
			tr.Expires >= 300 {
			continue
		}

		currentBid := tr.CurrentBid
		if currentBid <= 0 {
			currentBid = tr.StartingBid
		}

		tr.CurrentBid = currentBid
		tradesFilter.AuctionInfo = append(tradesFilter.AuctionInfo, tr)
	}

	return tradesFilter
}
