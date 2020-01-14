package bao

import (
	"conscience-backend/model"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
{
    "access_token": "NUYLS59JP1C8HKQUTT5JVA",
    "network_id": "fabric",
    "channel": "meetingchain",
    "contract": "faceid",
    "sync": true,
    "args": {
        "func_name": "FaceIDService.RegisterFaceID",
        "params": [
            {
                "source_type": "potograh",
                "source_hash": "xx-wwwww-xx",
                "algorithm": "sha256",
                "labels": [
                    "test"
                ],
                "metadata": {
                    "local": "bejing",
                    "name": "renjun"
                }
            }
        ]

*/
func TestBlockchainAPI_Invoke(t *testing.T) {
	api := fabricAPI{url: "http://127.0.0.1:8080", network: "fabric"}

	var args []interface{}
	args = append(args, &model.FaceID{
		ID:         "",
		SourceType: "image",
		SourceHash: "xx-abc-yy",
		Algorithm:  "hash256",
		Labels:     []string{"test"},
		Metadata:   nil,
		Timestamp:  0,
	})
	txID, response, err := api.Invoke("NUYLS59JP1C8HKQUTT5JVA", "meetingchain", "faceid", "FaceIDService.RegisterFaceID", args)
	assert.NoError(t, err)
	fmt.Println(txID)
	fmt.Println(response)
}

func TestBlockchainAPI_Query(t *testing.T) {
	api := fabricAPI{url: "http://127.0.0.1:8080", network: "fabric"}

	var args []interface{}
	response, err := api.Query("NUYLS59JP1C8HKQUTT5JVA", "meetingchain", "faceid", "FaceIDService.GetUser", args)
	assert.NoError(t, err)
	fmt.Println(response)
}
