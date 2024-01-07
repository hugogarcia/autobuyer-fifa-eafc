package watch

import (
	"fmt"

	"github.com/hugogarcia/fifa-auto-buyer/logger"
)

func ResetWatch() {
	//logger.LogMessage(nil, "----- RESETTING WATCH -------")

	list := getMostRecentTrades()
	if len(list.AuctionInfo) == 0 {
		logger.LogMessage(nil, "getMostRecentTrades returned empty")
		return
	}

	id := list.AuctionInfo[0].TradeId
	MakeWatch(id)
	err := StopWatching(fmt.Sprintf("%d", id))
	if err != nil {
		logger.LogMessage(nil, "Error reseting watch", err)
	} else {
		logger.LogMessage(nil, "RESET SUCCESS")
	}
}
