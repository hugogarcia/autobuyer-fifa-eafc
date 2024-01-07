package auth

import (
	"os"
)

type eaAuthResponse struct {
	Code string `json:"code"`
}

type auth struct {
	ClientVersion    int64          `json:"clientVersion"`
	Ds               string         `json:"ds"`
	GameSku          string         `json:"gameSku"`
	Identification   identification `json:"identification"`
	IsReadOnly       bool           `json:"isReadOnly"`
	Locale           string         `json:"locale"`
	Method           string         `json:"method"`
	NucleusPersonaID int64          `json:"nucleusPersonaId"`
	PriorityLevel    int64          `json:"priorityLevel"`
	Sku              string         `json:"sku"`
}

type identification struct {
	AuthCode    string `json:"authCode"`
	RedirectURL string `json:"redirectUrl"`
	Tid         string `json:"tid"`
}

var reqAuth = auth{
	ClientVersion:    1, //1 - PC, 3 - ANDROID
	Ds:               "5cd338759c2840e4a793053b12b6eb60aa930183cea8dec69e493d05de17ac33",
	GameSku:          os.Getenv("GAME_SKU"), //FFA24PS5
	IsReadOnly:       false,
	Locale:           "en-US",
	Method:           "authcode",
	NucleusPersonaID: 1280024159,
	PriorityLevel:    4,
	Sku:              "FUT24WEB",
	Identification: identification{
		AuthCode:    "",
		RedirectURL: "nucleus:rest",
	},
}

type AuthResponse struct {
	Protocol       string `json:"protocol"`
	IPPort         string `json:"ipPort"`
	ServerTime     string `json:"serverTime"`
	LastOnlineTime string `json:"lastOnlineTime"`
	Sid            string `json:"sid"`
	PhishingToken  string `json:"phishingToken"`
}

var currentSKU = 1
var clientsSkus = []auth{
	{ClientVersion: 3, Sku: "FUT24AND"},
	{ClientVersion: 1, Sku: "FUT24WEB"},
}
