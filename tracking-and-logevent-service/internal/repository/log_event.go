package repository

import "github.com/MamangPermen/tracking-service/internal/model"

// LogEventRepository menangani operasi penulisan log ke database NoSQL (Write-Heavy).
type LogEventRepository interface {
	InsertLog(log model.TrackingLog) error
}