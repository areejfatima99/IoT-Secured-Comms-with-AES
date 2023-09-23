package config

import (
	"custom-blockchain-iot-protocol/pkg"
	"errors"
	"os"
)

type Config struct {
	EthHost       string
	PrivateKey    string
	EncryptionKey string
}

func NewConfig() (*Config, error) {
	var config Config
	var ok bool
	if config.EthHost, ok = os.LookupEnv("ETH_HOST"); !ok {
		return nil, errors.New("ETH_HOST not set")
	}

	if config.PrivateKey, ok = os.LookupEnv("PRIVATE_KEY"); !ok {
		return nil, errors.New("PRIVATE_KEY not set")
	}

	if config.EncryptionKey, ok = os.LookupEnv("ENCRYPTION_KEY"); !ok {
		config.EncryptionKey = string(pkg.GenerateRandomKey())
	}

	return &config, nil
}
