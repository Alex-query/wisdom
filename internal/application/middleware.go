package application

import (
	"wisdom/internal/domain/service"
)

func middleWareCheck(challengeService service.ChallengeService, clientID string, req RequestMessageDTO) (bool, error) {
	ok, err := challengeService.VerifySolution(clientID, req.Meta.SecurityToken)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
