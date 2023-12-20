package service

type ChallengeService interface {
	GenerateTaskToResolve(clientID string) (string, error)
	VerifySolution(clientID string, solution string) (bool, error)
	Mint(prefixToken string) (string, error)
}
