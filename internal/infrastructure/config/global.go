package config

type GlobalConfigs struct {
	server    ServerConfig
	client    ClientConfig
	challenge ChallengeConfig
}

func GetGlobalConfigs() GlobalConfigs {
	globalConfigs := GlobalConfigs{
		server:    getServerConfig(),
		challenge: getChallengeConfig(),
		client:    getClientConfig(),
	}
	return globalConfigs
}

func (globalConfigs *GlobalConfigs) GetServerConfig() ServerConfig {
	return globalConfigs.server
}

func (globalConfigs *GlobalConfigs) GetChallengeConfig() ChallengeConfig {
	return globalConfigs.challenge
}

func (globalConfigs *GlobalConfigs) GetClientConfig() ClientConfig {
	return globalConfigs.client
}
