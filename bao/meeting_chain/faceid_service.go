package meeting_chain

import (
	"conscience-backend/bao"
	"conscience-backend/model"
	"fmt"
)

type FaceIDService struct {
	contract bao.IContract
}

func NewFaceIDService(contract bao.IContract) *FaceIDService {
	return &FaceIDService{contract: contract}
}

func (s *FaceIDService) RegisterFaceID(accessToken string, id *model.FaceID) (string, string, error) {
	var args []interface{}
	args = append(args, id)
	var tmpStr string
	txId, tmpStr, err := s.contract.Invoke(accessToken,
		"FaceIDService.RegisterFaceID",
		args,
	)
	fmt.Println(txId)
	return txId, tmpStr, err
}

func (s *FaceIDService) RegisterCertificate(accessToken string, id *model.FaceID) (string, string, error) {
	var args []interface{}
	args = append(args, id)
	var tmpStr string
	txId, tmpStr, err := s.contract.Invoke(accessToken,
		"FaceIDService.RegisterCertificate",
		args,
	)
	fmt.Println(txId)
	return txId, tmpStr, err
}

func (s *FaceIDService) Record(accessToken string, id *model.FaceID) (string, string, error) {
	var args []interface{}
	args = append(args, id)
	var tmpStr string
	txId, tmpStr, err := s.contract.Invoke(accessToken,
		"FaceIDService.Record",
		args,
	)
	fmt.Println(txId)
	return txId, tmpStr, err
}

func (s *FaceIDService) GetUser(accessToken string) (string, error) {
	var args []interface{}
	//args = append(args, id)
	txId, err := s.contract.Query(accessToken,
		"FaceIDService.GetUser",
		args,
	)
	fmt.Println(txId)
	return txId, err
}

func (s *FaceIDService) HistoryFaceIDs(accessToken string, id *model.HistoryFaceIDs) (string, error) {
	var args []interface{}
	args = append(args, id)
	txId, err := s.contract.Query(accessToken,
		"FaceIDService.HistoryFaceIDs",
		args,
	)
	fmt.Println(txId)
	return txId, err
}