package application

import (
	"encoding/json"
	"errors"
	"log"
	"time"
	"wisdom/internal/domain/entity"
	"wisdom/internal/domain/service"
)

type ApplicationServiceClient struct {
	client           service.ClientService
	errorChannel     chan error
	challengeService service.ChallengeService
	currentToken     string
}

func NewApplicationServiceClient(
	client service.ClientService,
	errorChannel chan error,
	challengeService service.ChallengeService,
) *ApplicationServiceClient {
	return &ApplicationServiceClient{
		client:           client,
		errorChannel:     errorChannel,
		challengeService: challengeService,
	}
}

func (service *ApplicationServiceClient) Init() error {
	readChannel := make(chan entity.ServerMessage)
	err := service.client.SubscribeMessages(readChannel, service.errorChannel)
	if err != nil {
		return err
	}
	go service.Listen(readChannel)
	return nil
}

func (service *ApplicationServiceClient) Listen(readChannel chan entity.ServerMessage) {
	for {
		message := <-readChannel
		log.Println("client message received ", string(message.Content))
		draftMessage := ResponseMessageBaseDTO{}
		err := json.Unmarshal(message.Content, &draftMessage)
		if err != nil {
			service.sendErrorMessage(err)
		}
		if draftMessage.Command == RequestMessageGetSecurityCodeDTOCommand {
			resp := ResponseMessageGetSecurityCodeDTO{}
			err := json.Unmarshal(message.Content, &resp)
			if err != nil {
				service.sendErrorMessage(err)
			}
			service.currentToken, err = service.challengeService.Mint(draftMessage.Meta.TaskToResolve)
			if err != nil {
				service.errorChannel <- err
			}
		}
		if draftMessage.Command == RequestMessageGetWisdomDTOCommand {
			resp := ResponseMessageGetWisdomDTO{}
			err := json.Unmarshal(message.Content, &resp)
			if err != nil {
				service.sendErrorMessage(err)
			}
			service.currentToken = ""
			log.Println("got wisdom: ", resp.Data.Wisdom)
			service.currentToken, err = service.challengeService.Mint(resp.Meta.TaskToResolve)
			if err != nil {
				service.errorChannel <- err
			}
		}
	}
}

func (service *ApplicationServiceClient) GetWisdom() error {
	if service.currentToken == "" {
		req := RequestMessageGetSecurityCodeDTO{}
		req.Command = RequestMessageGetSecurityCodeDTOCommand
		service.sendObjMessage(req)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		if service.currentToken != "" {
			break
		}
	}
	if service.currentToken == "" {
		return errors.New("security token is invalid")
	}
	req := RequestMessageGetWisdomDTO{}
	req.Command = RequestMessageGetWisdomDTOCommand
	req.Meta.SecurityToken = service.currentToken
	service.sendObjMessage(req)
	return nil
}

func (service *ApplicationServiceClient) sendErrorMessage(err error) {
	out := ErrorMessageDTO{
		ErrorMessage: err.Error(),
	}
	out.Meta.Code = 500
	service.sendObjMessage(out)
}

func (service *ApplicationServiceClient) sendObjMessage(obj interface{}) {
	message, err := json.Marshal(obj)
	if err != nil {
		service.errorChannel <- err
	}
	service.sendMessage(message)
}

func (service *ApplicationServiceClient) sendMessage(message []byte) {
	log.Println("client message sent ", string(message))
	err2 := service.client.SendMessage(entity.ServerMessage{
		Content: message,
	})
	if err2 != nil {
		service.errorChannel <- err2
	}
}
