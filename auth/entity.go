package auth

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
	Tid string `json:"tid"`
}

var reqAuth = auth{
	//ClientVersion: 3, //1 - PC, 3 - ANDROID
	ClientVersion: 1, //1 - PC, 3 - ANDROID
	Ds: "",
	GameSku: "FFA23PCC",
	IsReadOnly: false,
	Locale: "en-US",
	Method: "authcode",
	NucleusPersonaID: 1280024159,
	PriorityLevel: 4,
	Sku: "FUT23WEB", // FUT2WEB ou FUT23AND
	//Sku: "FUT23AND", // FUT2WEB ou FUT23AND
	Identification: identification{
		AuthCode: "",
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


var currentSKU = 0;
var clientsSkus = []auth{
	{ClientVersion: 3, Sku: "FUT23AND"},
	{ClientVersion: 1, Sku: "FUT23WEB"},
	{ClientVersion: 3, Sku: "FUT23IOS"},
}
