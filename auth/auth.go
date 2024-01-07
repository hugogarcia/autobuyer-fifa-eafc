package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	_ "github.com/joho/godotenv/autoload"
)

const (
	client_id       = "FUTWEB_BK_OL_SERVER"
	redirect_uri    = "nucleus:rest"
	response_type   = "code"
	release_type    = "prod"
	client_sequence = "ut-auth"
)

var access_token = os.Getenv("TOKEN")

var ea_URL = fmt.Sprintf("https://accounts.ea.com/connect/auth?client_id=%s&redirect_uri=%s&response_type=%s&access_token=%s&release_type=%s&client_sequence=%s",
	client_id, redirect_uri, response_type, access_token, release_type, client_sequence)

var authRequestCount = 0

func SetNewToken() {
	logger.LogMessage(nil, "------- SETTING NEW TOKEN -------")

	eaToken := getEAToken()
	if eaToken == "" {
		logger.PanicIt(nil, "EA token missing or empty")
	}

	url := "https://" + fifa.Host + "/ut/auth"

	fmt.Println(eaToken)
	reqAuth.Identification.AuthCode = eaToken
	reqAuth.NucleusPersonaID, _ = strconv.ParseInt(os.Getenv("USER_ID"), 10, 64)
	bd, _ := json.Marshal(reqAuth)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bd))
	if err != nil {
		logger.PanicIt(err, "Error requesting auth")
	}

	req.Header = fifa.FifaHeaders
	req.Header.Del("X-UT-SID")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.PanicIt(err, "ERROR ON REQUESTING AUTH")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.PanicIt(err, "ERROR GETTING BODY", ea_URL)
	}

	if resp.StatusCode != http.StatusOK {
		logger.PanicIt(err, fmt.Sprintf("IT WAS NOT POSSIBLE TO GET THE TOKEN FOR ULTIMATE TEAM. STATUS: %d \n %s", resp.StatusCode, string(body[:])))
	}

	var response AuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.PanicIt(err, string(body[:]))
	}

	logger.LogMessage(nil, "TOKEN ##############################", response.Sid)
	fifa.TOKEN_UT = response.Sid
	if fifa.TOKEN_UT == "" {
		logger.PanicIt(nil, "TOKEN EMPTY")
	}
	setNextAuthSku()
	authRequestCount++
	if authRequestCount >= 5 {
		logger.PanicIt(nil, "TOO MUCH TOKEN REQUEST, WAIT A MOMENT")
	}
}

func getEAToken() string {
	logger.LogMessage(nil, "GETTING EA AUTH")

	req, err := http.NewRequest(http.MethodGet, ea_URL, nil)
	if err != nil {
		logger.PanicIt(err, "SOMETHING WENT WRONT CREATING HTTP REQUEST \n"+err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(nil, "ERROR ON REQUESTING", ea_URL, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogMessage(nil, "ERROR GETTING BODY", ea_URL, err)
	}

	var response eaAuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.PanicIt(err, string(body[:]))
	}

	if response.Code == "" {
		logger.PanicIt(nil, "Error getting token:", ea_URL)
	}

	return response.Code
}

func setNextAuthSku() {
	currentSKU += 1
	if currentSKU > len(clientsSkus)-1 {
		currentSKU = 0
	}

	clientSku := clientsSkus[currentSKU]
	reqAuth.ClientVersion = clientSku.ClientVersion
	reqAuth.Sku = clientSku.Sku
	logger.LogMessage(nil, "CHANGING AUTH SKU")
}
