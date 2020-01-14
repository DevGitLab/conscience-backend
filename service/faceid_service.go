package service

import (
	"conscience-backend/bao/meeting_chain"
	"conscience-backend/model"
)

type FaceIDService struct {
}

func NewFaceIDService() *FaceIDService {
	return &FaceIDService{}
}

func (s *FaceIDService) RegisterFaceID(req *model.RequestRegisterFaceID) (string, string, error) {
	id := &model.FaceID{
		ID:         req.ID,
		SourceType: req.SourceType,
		SourceHash: req.SourceHash,
		Algorithm:  req.Algorithm,
		Labels:     req.Labels,
		Metadata:   req.Metadata,
		Timestamp:  req.Timestamp,
	}
	return meeting_chain.DefaultFaceIDService.RegisterFaceID(req.AccessToken, id)
}

func (s *FaceIDService) RegisterCertificate(req *model.RequestRegisterCertificate) (string, string, error) {
	id := &model.FaceID{
		ID:         req.ID,
		SourceType: req.SourceType,
		SourceHash: req.SourceHash,
		Algorithm:  req.Algorithm,
		Labels:     req.Labels,
		Metadata:   req.Metadata,
		Timestamp:  req.Timestamp,
	}
	return meeting_chain.DefaultFaceIDService.RegisterCertificate(req.AccessToken, id)
}

func (s *FaceIDService) Record(req *model.RequestRecord) (string, string, error) {
	id := &model.FaceID{
		ID:         req.ID,
		SourceType: req.SourceType,
		SourceHash: req.SourceHash,
		Algorithm:  req.Algorithm,
		Labels:     req.Labels,
		Metadata:   req.Metadata,
		Timestamp:  req.Timestamp,
	}
	return meeting_chain.DefaultFaceIDService.Record(req.AccessToken, id)
}

func (s *FaceIDService) GetUser(req *model.RequestGetUser) (string, error) {
	return meeting_chain.DefaultFaceIDService.GetUser(req.AccessToken)
}

func (s *FaceIDService) HistoryFaceIDs(req *model.RequestHistoryFaceIDs) (string, error) {
	id := &model.HistoryFaceIDs{
		ID:         req.ID,
		StartTime:  req.StartTime,
		EndTime: 	req.EndTime,
		Labels:     req.Labels,
		Timestamp:  req.Timestamp,
	}
	return meeting_chain.DefaultFaceIDService.HistoryFaceIDs(req.AccessToken, id)
}