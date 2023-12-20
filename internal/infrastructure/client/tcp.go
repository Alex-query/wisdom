package client

import (
	"net"
	"wisdom/internal/domain/entity"
	"wisdom/internal/infrastructure/config"
	"wisdom/internal/infrastructure/server"
)

type TCP struct {
	readChannel  chan entity.ServerMessage
	errorChannel chan error
	config       config.ClientConfig
	listener     server.Listener
}

func NewTCPClient(config config.ClientConfig) *TCP {
	s := &TCP{
		config: config,
	}
	return s
}

func (s *TCP) SubscribeMessages(readChannel chan entity.ServerMessage, errorChannel chan error) error {
	s.readChannel = readChannel
	s.errorChannel = errorChannel
	conn, err := net.Dial("tcp", s.config.GetHost()+":"+s.config.GetPort())
	if err != nil {
		return err
	}
	s.listener = server.NewListener(conn)
	go s.handleMessages()
	return nil
}

func (s *TCP) handleMessages() {
	for {
		message, err := s.listener.Read()
		if err != nil {
			s.errorChannel <- err
			break
		}
		s.readChannel <- entity.ServerMessage{
			Content: message,
		}
	}
}

func (s *TCP) SendMessage(message entity.ServerMessage) error {
	err := s.listener.Write(message.Content)
	if err != nil {
		return err
	}
	return nil
}
