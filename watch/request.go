package watch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hugogarcia/fifa-auto-buyer/auth"
	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/request"
)

func MakeWatch(tradeIds uint64) {

	url := "https://" + fifa.Host + "/ut/game/fc24/watchlist"

	//logger.LogMessage(nil, "WATCHING  tradeId:", tradeIds)

	ts := []watchRequest{{tradeIds}}
	requestBody := auction{AuctionInfo: ts}
	bd, _ := json.Marshal(requestBody)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bd))
	if err != nil {
		logger.PanicIt(err)
	}

	req.Header = fifa.FifaHeaders
	req.Header.Set("X-UT-SID", fifa.TOKEN_UT)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(nil, "ERROR ON REQUESTING", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogMessage(nil, "ERROR GETTING BODY", url, err)
	}

	if resp.StatusCode == 401 {
		auth.SetNewToken()
		MakeWatch(tradeIds)
	} else if resp.StatusCode > 400 {
		logger.LogMessage(nil, "ERROR MAKING WATCHING, STATUS CODE: ", resp.StatusCode, string(body[:]))

	} else {
		logger.LogMessage(nil, "Watch success")
	}
}

// StopWatching tradeIds separated by commas
func StopWatching(tradeIds string) error {
	if tradeIds == "" {
		return nil
	}
	url := fmt.Sprintf("https://%s/ut/game/fc24/watchlist?tradeId=%s", fifa.Host, tradeIds)

	logger.LogMessage(nil, "Deleting --- tradeId:", tradeIds)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header = fifa.FifaHeaders
	req.Header.Set("X-UT-SID", fifa.TOKEN_UT)
	//req.Proto = "HTTP/1.1"
	//req.ProtoMajor = 1
	//req.ProtoMinor = 1

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(nil, "ERROR ON REQUESTING", err)
		return err
	}

	if resp.StatusCode == 401 {
		auth.SetNewToken()
	} else if resp.StatusCode > 400 {
		logger.LogMessage(nil, "ERROR STOP WATCHING, STATUS CODE: ", resp.StatusCode)
	}

	return err
}

func GetWatchList() entity.TradeResponse {
	path := "/ut/game/fc24/watchlist"
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

	return trades
}

func getMostRecentTrades() entity.TradeResponse {
	path := "/ut/game/fc24/transfermarket"
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

	var tradesUnwatched entity.TradeResponse
	for _, v := range trades.AuctionInfo {
		if v.Watched || v.Expires < 6 {
			continue
		}
		tradesUnwatched.AuctionInfo = append(tradesUnwatched.AuctionInfo, v)
	}

	return trades
}
