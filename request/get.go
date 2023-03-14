package request

import (
	"io"
	"net/http"

	"github.com/hugogarcia/fifa-auto-buyer/auth"
	"github.com/hugogarcia/fifa-auto-buyer/fifa"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
)

func GetBodyByPath(path string) ([]byte, error) {
	url := "https://" + fifa.Host + path
	return getBody(url)
}

func GetBodyByURL(url string) ([]byte, error) {
	return getBody(url)
}

func getBody(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.PanicIt(err, "SOMETHING WENT WRONT CREATING HTTP REQUEST \n"+err.Error())
	}
	req.Header = fifa.FifaHeaders
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Header.Set("X-UT-SID", fifa.TOKEN_UT)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogMessage(nil, "ERROR ON REQUESTING", url, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogMessage(nil, "ERROR GETTING BODY", url, err)
	}
	
	if resp.StatusCode == 401 {
		auth.SetNewToken()
		return getBody(url)
	}

	if resp.StatusCode > 400 {
		logger.LogMessage(nil, "GetBody - Error requesting - status code: ", resp.StatusCode, "\n", string(body[:]))
	}

	return body, err
}
