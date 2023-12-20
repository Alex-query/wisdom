package config

import (
	"github.com/spf13/viper"
	"strconv"
)

type ChallengeConfig struct {
	numberOfZeroBits uint
	saltLength       uint
	extra            string
}

func NewChallengeConfig(numberOfZeroBits uint, saltLength uint, extra string) ChallengeConfig {
	return ChallengeConfig{
		numberOfZeroBits: numberOfZeroBits,
		saltLength:       saltLength,
		extra:            extra,
	}
}

func getChallengeConfig() ChallengeConfig {
	viper.BindEnv("WS_CHALLENGE_NUMBER_OF_ZERO_BITS")
	viper.SetDefault("WS_CHALLENGE_NUMBER_OF_ZERO_BITS", "20")
	viper.BindEnv("WS_CHALLENGE_SALT_LENGTH")
	viper.SetDefault("WS_CHALLENGE_SALT_LENGTH", "8")
	viper.BindEnv("WS_CHALLENGE_EXTRA")
	viper.SetDefault("WS_CHALLENGE_EXTRA", "extra")

	challengeConfig := ChallengeConfig{
		extra: viper.Get("WS_CHALLENGE_EXTRA").(string),
	}
	numberOfZeroBitsInt, err := strconv.Atoi(viper.Get("WS_CHALLENGE_NUMBER_OF_ZERO_BITS").(string))
	if err != nil {
		panic(err)
	}
	challengeConfig.numberOfZeroBits = uint(numberOfZeroBitsInt)
	saltLength, err := strconv.Atoi(viper.Get("WS_CHALLENGE_SALT_LENGTH").(string))
	if err != nil {
		panic(err)
	}
	challengeConfig.saltLength = uint(saltLength)
	return challengeConfig
}

func (challengeConfig *ChallengeConfig) GetNumberOfZeroBits() uint {
	return challengeConfig.numberOfZeroBits
}

func (challengeConfig *ChallengeConfig) GetSaltLength() uint {
	return challengeConfig.saltLength
}

func (challengeConfig *ChallengeConfig) GetExtra() string {
	return challengeConfig.extra
}
