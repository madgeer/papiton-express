package repository

import "github.com/MamangPermen/tracking-service/internal/model"

// TrackingRepository menangani operasi pembacaan histori pelacakan (Read-Heavy).
type TrackingRepository interface {
	GetResiHistory(resiID string) (*model.TrackingHistory, error)
}