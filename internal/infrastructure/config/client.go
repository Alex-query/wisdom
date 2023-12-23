package config

import (
	"github.com/spf13/viper"
	"time"
)

type ClientConfig struct {
	host                string
	port                string
	waitResponseTimeout time.Duration
}

func getClientConfig() ClientConfig {
	viper.BindEnv("WS_CLIENT_PORT")
	viper.SetDefault("WS_CLIENT_PORT", "1444")
	viper.BindEnv("WS_CLIENT_HOST")
	viper.SetDefault("WS_CLIENT_HOST", "127.0.0.1")
	viper.BindEnv("WS_CLIENT_WAIT_RESPONSE_TIMEOUT")
	viper.SetDefault("WS_CLIENT_WAIT_RESPONSE_TIMEOUT", "10s")

	waitResponseTimeout, err := time.ParseDuration(viper.Get("WS_CLIENT_WAIT_RESPONSE_TIMEOUT").(string))
	if err != nil {
		panic(err)
	}

	ClientConfig := ClientConfig{
		port:                viper.Get("WS_CLIENT_PORT").(string),
		host:                viper.Get("WS_CLIENT_HOST").(string),
		waitResponseTimeout: waitResponseTimeout,
	}
	return ClientConfig
}

func (ClientConfig *ClientConfig) GetPort() string {
	return ClientConfig.port
}

func (ClientConfig *ClientConfig) GetHost() string {
	return ClientConfig.host
}

func (ClientConfig *ClientConfig) GetWaitResponseTimeout() time.Duration {
	return ClientConfig.waitResponseTimeout
}
