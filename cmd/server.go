package cmd

import (
	"context"
	"log"
	"wisdom/internal/application"
	"wisdom/internal/infrastructure/challenge"
	"wisdom/internal/infrastructure/config"
	"wisdom/internal/infrastructure/repository"
	"wisdom/internal/infrastructure/server"
)

func RunServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	globalConfigs := config.GetGlobalConfigs()
	wisdomRepository := repository.NewWisdomRepository()
	hashier := challenge.NewHashierSha1(globalConfigs.GetChallengeConfig())
	challengeService := challenge.NewChallengeService(hashier)
	tcpServer := server.NewTCPServer(ctx, globalConfigs.GetServerConfig())
	errorChannel := make(chan error)
	go func() {
		for {
			err := <-errorChannel
			log.Println(err)
			//send to sentry
		}
	}()
	app := application.NewApplicationServiceServer(tcpServer, errorChannel, challengeService, wisdomRepository)
	err := app.Init()
	if err != nil {
		panic(err)
	}
}
