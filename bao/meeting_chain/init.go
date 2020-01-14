package meeting_chain

import "conscience-backend/bao"

var DefaultFaceIDService *FaceIDService

func Init(api bao.BlockchainAPI) {
	contract := NewContract(api)
	DefaultFaceIDService = NewFaceIDService(contract)
}
