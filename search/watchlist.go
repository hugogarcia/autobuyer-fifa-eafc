package search

import (
	"fmt"
	"strings"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/bid"
	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/watch"
)

func GetWatchlistAndRebid() {
	logger.LogMessage(nil, "Watching trades")

	trades := watch.GetWatchList()

	tradesToDelete := []uint64{}
	for _, v := range trades.AuctionInfo {
		nextBid := getNextBid(v.CurrentBid)

		player := entity.Players[v.ItemData.AssetId]
		if player.MaxBid <= 0 {
			logger.LogMessage(nil, fmt.Sprintf("Maxbid not found for playerId: %d -- Trade: %d", v.ItemData.AssetId, v.TradeId))
			tradesToDelete = append(tradesToDelete, v.TradeId)
			continue
		}

		//Ignore trades I'm already bid or time is too long
		if v.BidState == "highest" || v.Expires >= 60 {
			continue
		}

		//Delete trades when is not available
		if nextBid > player.MaxBid || v.Expires <= 0 || v.TradeState != "active" {
			tradesToDelete = append(tradesToDelete, v.TradeId)
			continue
		}

		t := GetTradeById(v.TradeId)
		nextBid = getNextBid(t.CurrentBid)

		if nextBid > player.MaxBid {
			tradesToDelete = append(tradesToDelete, v.TradeId)
			continue
		}

		logger.LogMessage(nil, "BIDDING----------------------")
		bid.MakeBid(v.TradeId, nextBid, player.Name)

		time.Sleep(time.Second * 1)
	}

	go deleteBids(tradesToDelete)

	time.Sleep(20 * time.Second)
	GetWatchlistAndRebid()
}

func deleteBids(trades []uint64) {
	if len(trades) == 0 {
		return
	}

	logger.LogMessage(nil, "Checking trades to unwatch...")
	ids := []string{}
	for _, v := range trades {
		ids = append(ids, fmt.Sprintf("%d", v))
		time.Sleep(time.Second)
	}

	watch.StopWatching(strings.Join(ids, ","))
}
