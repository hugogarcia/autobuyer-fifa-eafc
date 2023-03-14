package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
)

const (
	client_id       = "FUTWEB_BK_OL_SERVER"
	redirect_uri    = "nucleus:rest"
	response_type   = "code"
	release_type    = "prod"
	client_sequence = "ut-auth"
	access_token    = "eyJraWQiOiJkZGMxMDNkNC02ODQzLTRkNGQtYjhjMi04NTVmZDkyNjcyMWQiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJhY2NvdW50cy5lYS5jb20iLCJqdGkiOiJTa0V4T2pNdU1Eb3hMakE2Tm1Sak9XRXhabUl0WkdZd055MDBPREEwTFdJME1UVXRZalpqWWpZell6SXdOMlZqIiwiYXpwIjoiRklGQTIzX0pTX1dFQl9BUFAiLCJpYXQiOjE2Nzg4MTMxOTAsImV4cCI6MTY3ODgyNzU5MCwidmVyIjoxLCJuZXh1cyI6eyJyc3ZkIjp7ImVmcGx0eSI6IjEzIn0sImNsaSI6IkZJRkEyM19KU19XRUJfQVBQIiwicHJkIjoiMDFlNzU3NDItM2QyMy1jNzAzLTE0YWMtNTRjOWQzNDAwNmUyIiwic2NvIjoib2ZmbGluZSBzZWN1cml0eS50d29mYWN0b3Igc2lnbmluIGRwLmNvbW1lcmNlLmNsaWVudC5kZWZhdWx0IGRwLmlkZW50aXR5LmNsaWVudC5kZWZhdWx0IGRwLmZpcnN0cGFydHljb21tZXJjZS5jbGllbnQuZGVmYXVsdCIsInBpZCI6IjI4NDEwNjcwODYiLCJwdHkiOiJOVUNMRVVTIiwidWlkIjoiMjg0MTA2NzA4NiIsImR2aWQiOiI5MDg0YmQyNy01MjIxLTQ4ODctOGE0Zi0yOTA2Y2I1MmI0ODYiLCJwbHR5cCI6IldFQiIsInBuaWQiOiJFQSIsImRwaWQiOiJXRUIiLCJzdHBzIjoiT0ZGIiwidWRnIjpmYWxzZSwiY250eSI6IjEiLCJhdXNyYyI6IjMxNDY1NCIsImlwZ2VvIjp7ImlwIjoiMTMxLjEwMC4qLioiLCJjdHkiOiJCUiIsInJlZyI6IlBhcmFuYSIsImNpdCI6Ik1hcmluZ8OhIiwiaXNwIjoiTGlndWUgTW92ZWwgUy5BLiIsImxhdCI6Ii0yMy4zNzg3IiwibGd0IjoiLTUxLjk0ODkiLCJ0eiI6Ii0zIn0sInVpZiI6eyJ1ZGciOmZhbHNlLCJjdHkiOiJCUiIsImxhbiI6InB0Iiwic3RhIjoiQUNUSVZFIiwiYW5vIjpmYWxzZSwiYWdlIjozMX0sInBzaWYiOlt7ImlkIjoxMjgwMDI0MTU5LCJucyI6ImNlbV9lYV9pZCIsImRpcyI6Imh1Z29nYWFyY2lhIiwibmljIjoiaHVnb2dhYXJjaWEifV0sImVuYyI6IjByelhoNXRNRjIxUThFMUJnUHZWS1BOUlZaeHRtYjMraE1BMTZLK0xidk8yOXNpZllsYVVHcnB1UU80MU1vN1ZxOWdMQWptTVhsdzNKRTl3NHRkczhvWlhCWEkzWXlnRmQ4cDYrZUhVUy8rZ2h3WjlLclE3dm9pcXEvaGlBM1Z3ekVpakNZRkF0M3kvNXZGZFFKemttbzZ1c244aDlweFRHdUN5a2Jra1R3aHkyZ2o0bWJYTkVmVVhCMjhXMFE1SW9aOUozNVpQeDQ5alpQMVFENWtXdG8wdXI4Ync1ZDYxVkg0WHRpTy9FSDNwTkhGOW9UbmtubFdSWDQ3bWFCcXdNQnkyVEFDaUZnMktveWZpS0J2dHVteEpqdnJVdTd0bUxTTHNuK1Q4VHFGTkNaNGdOc0JuRVhrVHlRanhlQVFZYzFJY1pKSnBpY28wbmVhK3VVbHBvQTJRR2EyOFovTFdYd2hyckg1VHBNaERRRjlaTml0NVBRSEVFTDhvN1k5bkg2V2FzbjdKaEV4M1NFSHM2bEN5NXBzMFlsNE9zS3BRR2xrQ2ZaMXB1OUFyK1F3SVZ0SURsK05jaWZKZ2Q5SlNrYTR1d0ZoSVk0YTlTa1BlUGVsK01BSVNvZGovN1hGZmtpcC96T25xYUp1Q0cyenh1M0cwVjZpdit2amd0TTVZN05uRWxOdHd0amlNYjBXQ2tKa3VPUWJ4bERBbTN4bnVJWjVoQ2lLWEhvdVZDWVYxUUExU1pjaGk0MTJmQkczUllDa2FQRTlldHdCREJyMmcvNVdEdnI0MHg0dnJaaDEyQmlCUzcyTjZURXRlcGVWVkFFQTZpUWo0RzFmNU03ZTNQbkFjQ2FjK2c4a0JBalpXdjhxNmI1Z1Jya1lMNlNPa1dMb2VWa2Z0eUFNaGNsYkNJZ2U4ZDZLZHJKaGxmMGRpWlFzY1B5SjBXYU9CRm1LaFdzVGlHSzdCRzhmUkdpTUU4UjJzTEdjZ3VvTFM2MDBFT0pSU0FXZ3ZXaUJHTEpJMlMwREF4SVZIZEtLTzhncTZiN1NPTjY2MWRZRllLVUpoVzVqODB4K0RKS1laaDdpTWdXc2t3OXk1bWRVcURaaXF1NURxUk13MlltNkpFY0I3Yk81dUNCTldYdEcycFZuNnpVUEpTUWdOVkNVMjNYMzlLU0ozWmd4dTAxbms0ejZhRmQyN3Brdjl3dm1MR2lSRWI0cVdremZUTGduN0Rab2F4NlhUQXZwNFJDWE8wakQxRGdoS0ZhL2dFN2NFcmJ4Ujg2Q29XNEhOcUtBZ0JTejl4bHB1Zjc1TjNWRG5JdTRFSW1idlRQek9RYjA9In19.S8C1dpeWNpHrwKSVHY1gBWP9zy7MK_figYEjKut8CgTU4JhdnELWlSEW2eSZVzEpOLLAIKWr1b9avBfLC30yYDRoNKfBgTjSr9g3PNJIdw0RqnIc3pMeuZB7t8dsOuIcoO2Q3U2a_-NwOJ9wmNuyEjfRBlxXrZVxh9veGtWPcIyWO2Yu2_zelUk8qRRZo3SVwjpFyB-fuK17vSTGcdXE5g32HWdSTHwVQO6m1QeKdlGdxtwWV_lMarEIRjXjey_fOZQ40fz04w6DALM-MlqM8wc5fY3VpOHOGESOaD99pUdSKpT9rCCbLycIBvXnhCym4G0sp2mIhs7mbNcyPIG00w"
)

var ea_URL = fmt.Sprintf("https://accounts.ea.com/connect/auth?client_id=%s&redirect_uri=%s&response_type=%s&access_token=%s&release_type=%s&client_sequence=%s",
	client_id, redirect_uri, response_type, access_token, release_type, client_sequence)

var authRequestCount = 0

func SetNewToken() {
	logger.LogMessage(nil, "------- SETTING NEW TOKEN -------")

	ea_token := getEAToken()
	if ea_token == "" {
		logger.PanicIt(nil, "EA token missing or empty")
	}

	url := "https://" + fifa.Host + "/ut/auth"

	reqAuth.Identification.AuthCode = ea_token
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
	if authRequestCount >= 5{
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
	if currentSKU > len(clientsSkus) - 1 {
		currentSKU = 0
	}

	clientSku := clientsSkus[currentSKU]
	reqAuth.ClientVersion = clientSku.ClientVersion
	reqAuth.Sku = clientSku.Sku
	logger.LogMessage(nil, "CHANGING AUTH SKU")
}
