package service

import "wisdom/internal/domain/entity"

type ClientService interface {
	SubscribeMessages(readChannel chan entity.ServerMessage, errorChannel chan error) error
	SendMessage(message entity.ServerMessage) error
}
