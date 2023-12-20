package cmd

import (
	"log"
	"time"
	"wisdom/internal/application"
	"wisdom/internal/infrastructure/challenge"
	"wisdom/internal/infrastructure/client"
	"wisdom/internal/infrastructure/config"
)

func RunClient() {
	globalConfigs := config.GetGlobalConfigs()
	hashier := challenge.NewHashierSha1(globalConfigs.GetChallengeConfig())
	challengeService := challenge.NewChallengeService(hashier)
	tcpClient := client.NewTCPClient(globalConfigs.GetClientConfig())
	errorChannel := make(chan error)
	go func() {
		for {
			err := <-errorChannel
			log.Println(err)
			//send to sentry
		}
	}()
	app := application.NewApplicationServiceClient(tcpClient, errorChannel, challengeService)
	err := app.Init()
	if err != nil {
		panic(err)
	}
	err = app.GetWisdom()
	if err != nil {
		panic(err)
	}
	time.Sleep(133 * time.Second)
}
