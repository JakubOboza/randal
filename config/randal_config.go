package config

import (
	"fmt"
	"math/rand"
	"net/url"

	"gopkg.in/yaml.v3"
)

type RandalConfig struct {
	RootUrl   string                     `yaml:"root_url" json:"root_url" `
	EndPoints map[string]*RandalEndpoint `yaml:"endpoints" json:"endpoints"`
}

type RandalEndpoint struct {
	Destinations []string `yml:"destinations" json:"destinations"`
}

func Load(rawInput []byte) (*RandalConfig, error) {
	cfg := &RandalConfig{}
	err := yaml.Unmarshal(rawInput, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (re RandalEndpoint) IsValid() bool {
	valid := true
	if len(re.Destinations) < 1 {
		valid = false
	}
	for _, candidate := range re.Destinations {
		_, err := url.Parse(candidate)
		if err != nil {
			fmt.Println(err)
			fmt.Println("invalid url: ", candidate)
			valid = false
		}
	}
	return valid
}

func (re RandalEndpoint) Next() string {
	destSize := len(re.Destinations)
	if destSize > 0 {
		i := rand.Intn(destSize)
		return re.Destinations[i]
	}
	return ""
}
