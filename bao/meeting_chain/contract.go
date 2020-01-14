package meeting_chain

import "conscience-backend/bao"

type Contract struct {
	channel  string
	contract string
	api      bao.BlockchainAPI
}

func (c *Contract) Invoke(accessToken, fn string, args []interface{}) (string, string, error) {
	return c.api.Invoke(accessToken, c.channel, c.contract, fn, args)
}

func (c *Contract) Query(accessToken, fn string, args []interface{}) (string, error) {
	return c.api.Query(accessToken, c.channel, c.contract, fn, args)
}

func NewContract(api bao.BlockchainAPI) bao.IContract {
	return &Contract{
		channel:  "meetingchain",
		contract: "faceid",
		api:      api,
	}
}
