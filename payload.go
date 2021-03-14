package go5paisa

type payloadHead struct {
	AppName     string `json:"appName"`
	AppVer      string `json:"appVer"`
	Key         string `json:"key"`
	OsName      string `json:"osName"`
	RequestCode string `json:"requestCode"`
	UserID      string `json:"userId"`
	Password    string `json:"password"`
}

type loginBody struct {
	Email          string `json:"Email_id"`
	Password       string `json:"Password"`
	LocalIP        string `json:"LocalIP"`
	PublicIP       string `json:"PublicIP"`
	SerialNumber   string `json:"HDSerailNumber"`
	MAC            string `json:"MACAddress"`
	MachineID      string `json:"MachineID"`
	VersionNo      string `json:"VersionNo"`
	RequestNo      string `json:"RequestNo"`
	My2PIN         string `json:"My2PIN"`
	ConnectionType string `json:"ConnectionType"`
}

type loginPayload struct {
	Head *payloadHead `json:"head"`
	Body loginBody    `json:"body"`
}

type genericPayload struct {
	Head *payloadHead       `json:"head"`
	Body genericPayloadBody `json:"body"`
}

type genericPayloadBody struct {
	ClientCode string `json:"ClientCode"`
}

type orderStatusPayloadBody struct {
	ClientCode string           `json:"ClientCode"`
	OrdList    []OrderForStatus `json:"OrdStatusReqList"`
}

type orderStatusPayload struct {
	Head *payloadHead           `json:"head"`
	Body orderStatusPayloadBody `json:"body"`
}

type marketFeedPayload struct {
	Head *payloadHead          `json:"head"`
	Body marketFeedPayloadBody `json:"body"`
}

type marketFeedPayloadBody struct {
	ClientCode      string            `json:"ClientCode"`
	ClientLoginType int               `json:"ClientLoginType"`
	MarketFeedData  []MarketFeedQuery `json:"MarketFeedData"`
}
