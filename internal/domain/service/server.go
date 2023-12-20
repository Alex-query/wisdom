package service

import "wisdom/internal/domain/entity"

type ServerService interface {
	ServeAndListen(readChannel chan entity.ServerMessage, errorChannel chan error) error
	SendMessage(message entity.ServerMessage) error
}
