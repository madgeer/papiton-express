package service

import (
	"github.com/MamangPermen/tracking-service/internal/model"
	"github.com/MamangPermen/tracking-service/internal/repository"
)

type TrackingService struct {
	repo repository.TrackingRepository
}

func NewTrackingService(repo repository.TrackingRepository) *TrackingService {
	return &TrackingService{repo: repo}
}

func (s *TrackingService) GetHistory(resiID string) (*model.TrackingHistory, error) {
	// Memanggil repo sesungguhnya
	return s.repo.GetResiHistory(resiID)
}
