package service

import "wisdom/internal/infrastructure/config"

type SyncService interface {
	Init(clientConfig config.ClientConfig)
	GenerateRequestID() (string, error)
	WaitResponseByRequestID(requestID string) ([]byte, error)
	PushResponse(requestID string, response []byte, err error) error
}
