package challenge

import (
	"log"
	"sync"
	"wisdom/internal/domain/service"
)

type Hasher interface {
	Mint(prefixToken string) (string, error)
	GenerateRandomPrefixToken() (string, error)
	Check(stamp string, rightPrefix string) bool
}

type Service struct {
	Hasher        Hasher
	mapClientTask sync.Map // map[clientID]task
}

var _ service.ChallengeService = &Service{}

func NewChallengeService(
	Hasher Hasher,
) *Service {
	return &Service{
		Hasher: Hasher,
	}
}

func (c *Service) GenerateTaskToResolve(clientID string) (string, error) {
	task, err := c.Hasher.GenerateRandomPrefixToken()
	if err != nil {
		return "", err
	}
	c.mapClientTask.Store(clientID, task)
	return task, nil
}

func (c *Service) VerifySolution(clientID string, solution string) (bool, error) {
	task, ok := c.mapClientTask.Load(clientID)
	if !ok {
		log.Println("clientID not found")
		return false, nil
	}
	ok = c.Hasher.Check(solution, task.(string))
	return ok, nil
}

func (c *Service) Mint(prefixToken string) (string, error) {
	return c.Hasher.Mint(prefixToken)
}
