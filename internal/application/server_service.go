package application

import (
	"encoding/json"
	"errors"
	"log"
	"wisdom/internal/domain/entity"
	"wisdom/internal/domain/repository"
	"wisdom/internal/domain/service"
)

type ApplicationServiceServer struct {
	server           service.ServerService
	errorChannel     chan error
	challengeService service.ChallengeService
	wisdomRepository repository.WisdomRepository
}

func NewApplicationServiceServer(
	server service.ServerService,
	errorChannel chan error,
	challengeService service.ChallengeService,
	wisdomRepository repository.WisdomRepository,
) *ApplicationServiceServer {
	return &ApplicationServiceServer{
		server:           server,
		errorChannel:     errorChannel,
		challengeService: challengeService,
		wisdomRepository: wisdomRepository,
	}
}

func (service *ApplicationServiceServer) Init() error {
	readChannel := make(chan entity.ServerMessage)
	err := service.server.ServeAndListen(readChannel, service.errorChannel)
	if err != nil {
		return err
	}
	for {
		message := <-readChannel
		log.Println("server message received ", string(message.Content))
		draftMessage := RequestMessageDTO{}
		err := json.Unmarshal(message.Content, &draftMessage)
		if err != nil {
			service.sendErrorMessage(message.ClientID, err)
		}
		if draftMessage.Command != RequestMessageGetSecurityCodeDTOCommand {
			ok, err := middleWareCheck(service.challengeService, message.ClientID, draftMessage)
			if err != nil {
				service.errorChannel <- err
			}
			if !ok {
				service.sendErrorMessage(message.ClientID, errors.New("security token is invalid"))
				continue
			}
		}
		switch draftMessage.Command {
		case RequestMessageGetSecurityCodeDTOCommand:
			err = service.GetSecurityCode(message.ClientID, draftMessage)
			if err != nil {
				service.sendErrorMessage(message.ClientID, err)
			}
		case RequestMessageGetWisdomDTOCommand:
			err = service.GetWisdom(message.ClientID, draftMessage)
			if err != nil {
				service.sendErrorMessage(message.ClientID, err)
			}
		default:
			service.sendErrorMessage(message.ClientID, errors.New("unknown command"))
		}
	}
}

func (service *ApplicationServiceServer) GetSecurityCode(clientID string, request RequestMessageDTO) error {
	task, err := service.challengeService.GenerateTaskToResolve(clientID)
	if err != nil {
		return err
	}
	dto := ResponseMessageGetSecurityCodeDTO{}
	dto.Meta.Code = 200
	dto.Meta.TaskToResolve = task
	dto.Meta.RequestID = request.Meta.RequestID
	dto.Command = RequestMessageGetSecurityCodeDTOCommand
	service.sendObjMessage(clientID, dto)
	return nil
}

func (service *ApplicationServiceServer) GetWisdom(clientID string, request RequestMessageDTO) error {
	wisdom, err := service.wisdomRepository.GetRandomWisdom()
	if err != nil {
		return err
	}
	dto := ResponseMessageGetWisdomDTO{
		Data: ResponseMessageGetWisdomDTOData{
			Wisdom: wisdom,
		},
	}
	dto.Meta.TaskToResolve, err = service.challengeService.GenerateTaskToResolve(clientID)
	if err != nil {
		return err
	}
	dto.Meta.Code = 200
	dto.Meta.RequestID = request.Meta.RequestID
	dto.Command = RequestMessageGetWisdomDTOCommand
	service.sendObjMessage(clientID, dto)
	return nil
}

func (service *ApplicationServiceServer) sendErrorMessage(clientID string, err error) {
	out := ErrorMessageDTO{
		ErrorMessage: err.Error(),
	}
	out.Meta.Code = 500
	service.sendObjMessage(clientID, out)
}

func (service *ApplicationServiceServer) sendObjMessage(clientID string, obj interface{}) {
	message, err := json.Marshal(obj)
	if err != nil {
		service.errorChannel <- err
	}
	service.sendMessage(clientID, message)
}

func (service *ApplicationServiceServer) sendMessage(clientID string, message []byte) {
	log.Println("sending message to ", clientID, " message: ", string(message))
	err2 := service.server.SendMessage(entity.ServerMessage{
		ClientID: clientID,
		Content:  message,
	})
	if err2 != nil {
		service.errorChannel <- err2
	}
}
