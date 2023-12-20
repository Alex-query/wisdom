package server

import (
	"context"
	"errors"
	"net"
	"wisdom/internal/domain/entity"
	"wisdom/internal/domain/service"
	"wisdom/internal/infrastructure/config"
)

type TCP struct {
	listeners    []Listener
	readChannel  chan entity.ServerMessage
	errorChannel chan error
	config       config.ServerConfig
	ctx          context.Context
}

var _ service.ServerService = &TCP{}

func NewTCPServer(ctx context.Context, config config.ServerConfig) *TCP {
	s := &TCP{
		listeners: []Listener{},
		config:    config,
		ctx:       ctx,
	}
	return s
}

func (s *TCP) ServeAndListen(readChannel chan entity.ServerMessage, errorChannel chan error) error {
	s.readChannel = readChannel
	s.errorChannel = errorChannel
	listenerTCP, err := net.Listen("tcp", s.config.GetHost()+":"+s.config.GetPort())
	if err != nil {
		return err
	}
	go func() {
		<-s.ctx.Done()
		err := listenerTCP.Close()
		if err != nil {
			s.errorChannel <- err
		}
	}()
	go func() {
		for {
			conn, err := listenerTCP.Accept()
			if err != nil {
				s.errorChannel <- err
			}
			listener := NewListener(conn)
			s.listeners = append(s.listeners, listener)
			go s.handleConnection(listener)
		}
	}()
	return nil
}

func (s *TCP) SendMessage(message entity.ServerMessage) error {
	for _, listener := range s.listeners {
		if listener.GetClientID() == message.ClientID {
			err := listener.Write(message.Content)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("client not found")
}

func (s *TCP) handleConnection(listener Listener) {
	for {
		message, err := listener.Read()
		if err != nil {
			s.errorChannel <- err
			break
		}
		s.readChannel <- entity.ServerMessage{
			ClientID: listener.GetClientID(),
			Content:  message,
		}
	}
}
