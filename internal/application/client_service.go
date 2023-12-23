package application

import (
	"encoding/json"
	"log"
	"wisdom/internal/domain/entity"
	"wisdom/internal/domain/service"
)

type ApplicationServiceClient struct {
	client           service.ClientService
	errorChannel     chan error
	challengeService service.ChallengeService
	currentToken     string
	syncService      service.SyncService
}

func NewApplicationServiceClient(
	client service.ClientService,
	errorChannel chan error,
	challengeService service.ChallengeService,
	syncService service.SyncService,
) *ApplicationServiceClient {
	return &ApplicationServiceClient{
		client:           client,
		errorChannel:     errorChannel,
		challengeService: challengeService,
		syncService:      syncService,
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
		err = service.syncService.PushResponse(draftMessage.Meta.RequestID, message.Content, nil)
		if err != nil {
			service.errorChannel <- err
		}
	}
}

func (service *ApplicationServiceClient) reqSecurityToken() error {
	sr, err := service.syncService.GenerateRequestID()
	if err != nil {
		return err
	}
	req := RequestMessageGetSecurityCodeDTO{}
	req.Command = RequestMessageGetSecurityCodeDTOCommand
	req.Meta.RequestID = sr
	service.sendObjMessage(req)
	respText, err := service.syncService.WaitResponseByRequestID(sr)
	if err != nil {
		return err
	}
	resp := ResponseMessageGetSecurityCodeDTO{}
	err = json.Unmarshal(respText, &resp)
	if err != nil {
		service.sendErrorMessage(err)
	}
	service.currentToken, err = service.challengeService.Mint(resp.Meta.TaskToResolve)
	if err != nil {
		service.errorChannel <- err
	}
	return nil
}

func (service *ApplicationServiceClient) GetWisdom() (string, error) {
	if service.currentToken == "" {
		err := service.reqSecurityToken()
		if err != nil {
			return "", err
		}
	}
	sr, err := service.syncService.GenerateRequestID()
	if err != nil {
		return "", err
	}
	req := RequestMessageGetWisdomDTO{}
	req.Command = RequestMessageGetWisdomDTOCommand
	req.Meta.SecurityToken = service.currentToken
	req.Meta.RequestID = sr
	service.sendObjMessage(req)
	respText, err := service.syncService.WaitResponseByRequestID(sr)
	if err != nil {
		return "", err
	}
	resp := ResponseMessageGetWisdomDTO{}
	err = json.Unmarshal(respText, &resp)
	if err != nil {
		service.sendErrorMessage(err)
	}
	service.currentToken = ""
	service.currentToken, err = service.challengeService.Mint(resp.Meta.TaskToResolve)
	if err != nil {
		service.errorChannel <- err
	}
	return resp.Data.Wisdom, nil
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
