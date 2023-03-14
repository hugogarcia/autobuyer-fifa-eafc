package bid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/auth"
	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/watch"
)

var lastBid = time.Now()

var errorCount int;

func MakeBid(tradeId, nextBid uint64) error {
	url := fmt.Sprintf("https://"+fifa.Host+"/ut/game/fifa23/trade/%d/bid", tradeId)

	logger.LogMessage(nil, "------------------- MAKING BID --- tradeId:", tradeId, "--- bid:", nextBid)

	if time.Since(lastBid).Seconds() <= 2{
		logger.LogMessage(nil, "------------------- Waiting for bid...")
		time.Sleep(time.Second)
	}

	b := struct {
		Bid uint64 `json:"bid"`
	}{nextBid}

	bd, _ := json.Marshal(b)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bd))
	if err != nil {
		return err
	}
	
	req.Header = fifa.FifaHeaders
	req.Header.Set("X-UT-SID", fifa.TOKEN_UT)	
	req.Proto = "HTTP/1.1"

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(err, "----------------- ERROR ON REQUESTING", err)
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogMessage(err, "----------------- ERROR GETTING BODY", url, err)
	}

	if resp.StatusCode == http.StatusUnauthorized{
		auth.SetNewToken();
		return nil;
	}

	if resp.StatusCode >= 400 {
		logger.LogMessage(nil, "------------------- ERROR WHEN MAKING BID, STATUS CODE: ", resp.StatusCode, string(body))
		auth.SetNewToken();
		watch.ResetWatch()		
	} else {
		logger.LogMessage(nil, "--------------------------------- BID SUCCESS")
		errorCount = 0
	}

	lastBid = time.Now()

	return err
}

