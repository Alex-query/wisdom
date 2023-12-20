package config

import "github.com/spf13/viper"

type ServerConfig struct {
	host string
	port string
}

func getServerConfig() ServerConfig {
	viper.BindEnv("WS_SERVER_PORT")
	viper.SetDefault("WS_SERVER_PORT", "1444")
	viper.BindEnv("WS_SERVER_HOST")
	viper.SetDefault("WS_SERVER_HOST", "127.0.0.1")

	serverConfig := ServerConfig{
		port: viper.Get("WS_SERVER_PORT").(string),
		host: viper.Get("WS_SERVER_HOST").(string),
	}
	return serverConfig
}

func (serverConfig *ServerConfig) GetPort() string {
	return serverConfig.port
}

func (serverConfig *ServerConfig) GetHost() string {
	return serverConfig.host
}
