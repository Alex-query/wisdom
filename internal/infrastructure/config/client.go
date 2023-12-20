package config

import "github.com/spf13/viper"

type ClientConfig struct {
	host string
	port string
}

func getClientConfig() ClientConfig {
	viper.BindEnv("WS_CLIENT_PORT")
	viper.SetDefault("WS_CLIENT_PORT", "1444")
	viper.BindEnv("WS_CLIENT_HOST")
	viper.SetDefault("WS_CLIENT_HOST", "127.0.0.1")

	ClientConfig := ClientConfig{
		port: viper.Get("WS_CLIENT_PORT").(string),
		host: viper.Get("WS_CLIENT_HOST").(string),
	}
	return ClientConfig
}

func (ClientConfig *ClientConfig) GetPort() string {
	return ClientConfig.port
}

func (ClientConfig *ClientConfig) GetHost() string {
	return ClientConfig.host
}
