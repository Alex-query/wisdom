package cmd

import (
	"log"
	"wisdom/internal/application"
	"wisdom/internal/infrastructure/challenge"
	"wisdom/internal/infrastructure/client"
	"wisdom/internal/infrastructure/config"
	"wisdom/internal/infrastructure/sync"
)

func RunClient() {
	globalConfigs := config.GetGlobalConfigs()
	hashier := challenge.NewHashierSha1(globalConfigs.GetChallengeConfig())
	challengeService := challenge.NewChallengeService(hashier)
	tcpClient := client.NewTCPClient(globalConfigs.GetClientConfig())
	syncService := sync.NewServiceSync(globalConfigs.GetClientConfig())
	errorChannel := make(chan error)
	go func() {
		for {
			err := <-errorChannel
			log.Println(err)
			//send to sentry
		}
	}()
	app := application.NewApplicationServiceClient(
		tcpClient,
		errorChannel,
		challengeService,
		syncService,
	)
	err := app.Init()
	if err != nil {
		panic(err)
	}
	wis, err := app.GetWisdom()
	if err != nil {
		panic(err)
	}
	log.Println(wis)
	wis, err = app.GetWisdom()
	if err != nil {
		panic(err)
	}
	log.Println(wis)
}
