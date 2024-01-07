package watch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
)

type PinRequest struct {
	Custom Custom   `json:"custom"`
	Et     string   `json:"et"`
	Events []Events `json:"events"`
	Gid    int      `json:"gid"`
	IsSess bool     `json:"is_sess"`
	Loc    string   `json:"loc"`
	Plat   string   `json:"plat"`
	Rel    string   `json:"rel"`
	Sid    string   `json:"sid"`
	Taxv   string   `json:"taxv"`
	Tid    string   `json:"tid"`
	Tidt   string   `json:"tidt"`
	TsPost string   `json:"ts_post"`
	V      string   `json:"v"`
}
type Custom struct {
	NetworkAccess string `json:"networkAccess"`
	ServicePlat   string `json:"service_plat"`
}
type Pidm struct {
	Nucleus int64 `json:"nucleus"`
}
type Core struct {
	S       int    `json:"s"`
	Pidt    string `json:"pidt"`
	Pid     string `json:"pid"`
	TsEvent string `json:"ts_event"`
	En      string `json:"en"`
	Pidm    Pidm   `json:"pidm"`
}
type Events struct {
	Type string `json:"type"`
	Pgid string `json:"pgid"`
	Core Core   `json:"core"`
}

var data PinRequest //= getFromFile()

func getFromFile() PinRequest {
	content, err := os.ReadFile("event.json")
	if err != nil {
		logger.PanicIt(err)
	}

	var request PinRequest
	err = json.Unmarshal(content, &request)
	if err != nil {
		logger.PanicIt(err)
	}

	return request
}

func MakePin() {
	dt := time.Now().Add(3 * time.Hour).UTC().Format("2006-01-02T15:04:05")
	data.Sid = fifa.TOKEN_UT
	data.TsPost = dt
	for _, v := range data.Events {
		v.Core.TsEvent = dt
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logger.LogMessage(nil, err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, "https://pin-river.data.ea.com/pinEvents", body)
	if err != nil {
		logger.LogMessage(nil, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(nil, err)
	}
	defer resp.Body.Close()
}
