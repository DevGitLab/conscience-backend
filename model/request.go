package model

type RequestInvokeOrQueryContractV2 struct {
	AccessToken string `json:"access_token"`
	NetworkId   string `json:"network_id" ` // network id
	Channel     string `json:"channel" `    // channel id
	Contract    string `json:"contract" `   // contract name
	Sync        bool   `json:"sync" `       // true: means wait block
	Args        *Args  `json:"args" `
}

type Args struct {
	Method string        `json:"func_name" `
	Params []interface{} `json:"params" `
}

type RequestRegisterFaceID struct {
	AccessToken string `json:"access_token"`
	FaceID
}

type RequestRegisterCertificate struct {
	AccessToken string `json:"access_token"`
	FaceID
}

type RequestRecord struct {
	AccessToken string `json:"access_token"`
	FaceID
}

type RequestGetUser struct {
	AccessToken string `json:"access_token"`
}

type RequestHistoryFaceIDs struct {
	AccessToken string `json:"access_token"`
	HistoryFaceIDs
}