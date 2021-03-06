package config

import (
	"errors"
)

func RegisterConfigurationPartParser(key string, parser ConfigurationPartParser) error {
	if _, ok := configurationPartParsers[key]; ok {
		return errors.New("config: can't register config parser " + key + " it's been registered already")
	}
	configurationPartParsers[key] = parser
	return nil
}

func GetConfigurationPartParser(key string) (ConfigurationPartParser, error) {
	if p, ok := configurationPartParsers[key]; ok {
		return p, nil
	}
	return nil, errors.New("Unrecognized configuration key: " + key)
}

var configurationPartParsers = map[string]ConfigurationPartParser{}
