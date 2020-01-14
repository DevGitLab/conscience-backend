package meeting_chain

import (
	"conscience-backend/bao"
	"conscience-backend/config"
	"conscience-backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFaceIDService_RegisterFaceID(t *testing.T) {
	api := bao.NewBlockchainAPI(config.BlockchainAPI{Url: "http://127.0.0.1:8080"})
	contract := NewContract(api)
	faceIDService := FaceIDService{contract: contract}
	id := &model.FaceID{
		ID:         "",
		SourceType: "image",
		SourceHash: "xx-abc-yy",
		Algorithm:  "hash256",
		Labels:     []string{"test"},
		Metadata:   nil,
		Timestamp:  0,
	}
	err := faceIDService.RegisterFaceID("HGN4HCRENAWKPTSLH-8UDW", id)
	assert.NoError(t, err)
}
