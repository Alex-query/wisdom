package sync

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
	"wisdom/internal/domain/service"
	"wisdom/internal/infrastructure/config"
)

type ServiceSync struct {
	requests     sync.Map
	clientConfig config.ClientConfig
}

var _ service.SyncService = &ServiceSync{}

func NewServiceSync(clientConfig config.ClientConfig) *ServiceSync {
	return &ServiceSync{
		clientConfig: clientConfig,
	}
}

func (s *ServiceSync) Init(clientConfig config.ClientConfig) {
	s.clientConfig = clientConfig
	s.requests = sync.Map{}
}

func (s *ServiceSync) GenerateRequestID() (string, error) {
	uuidS := uuid.New().String()
	respChan := make(chan []byte)
	errChan := make(chan error)
	s.requests.Store(uuidS, Request{
		ID:         uuidS,
		chResponse: respChan,
		chError:    errChan,
		createdAt:  time.Now(),
	})
	return uuidS, nil
}

func (s *ServiceSync) PushResponse(requestID string, response []byte, err error) error {
	req, ok := s.requests.Load(requestID)
	if !ok {
		return errors.New("request not found")
	}
	request := req.(Request)
	if err != nil {
		request.chError <- err
		return nil
	}
	request.chResponse <- response
	return nil
}

func (s *ServiceSync) WaitResponseByRequestID(requestID string) ([]byte, error) {
	req, ok := s.requests.Load(requestID)
	if !ok {
		return []byte{}, errors.New("request not found")
	}
	request := req.(Request)
	select {
	case resp := <-request.chResponse:
		return resp, nil
	case err := <-request.chError:
		return []byte{}, err
	case <-time.After(s.clientConfig.GetWaitResponseTimeout()):
		return []byte{}, errors.New("timeout")
	}
	return []byte{}, nil
}
