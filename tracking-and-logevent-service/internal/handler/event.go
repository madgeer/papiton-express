package handler

import (
	"encoding/json"
	"github.com/MamangPermen/tracking-service/internal/model"
	"github.com/MamangPermen/tracking-service/internal/service"
)

type EventConsumerHandler struct {
	svc *service.LogEventService
}

func NewEventConsumerHandler(svc *service.LogEventService) *EventConsumerHandler {
	return &EventConsumerHandler{svc: svc}
}

func (h *EventConsumerHandler) ConsumeLogEvent(message []byte) {
	var log model.TrackingLog
	if err := json.Unmarshal(message, &log); err != nil {
		return // Silently fail atau log error untuk background process
	}
	
	_ = h.svc.ProcessLog(log)
}
