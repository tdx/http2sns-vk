package config

import (
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

//
type Config struct {
	HttpListenAddr    string         `envconfig:"HTTP_LISTEN_ADDR" default:":80"`
	HttpEndpointTopic MapEndpointArn `envconfig:"HTTP_ENDPOINT_TOPIC"`
	HttpDebug         bool           `envconfig:"HTTP_DEBUG"       default:"false"`
	SnsApiEndpoint    string         `envconfig:"SNS_API_ENDPOINT"`
	SnsRegion         string         `envconfig:"REGION"           default:"eu-central-1"`
}

// NewConfig creates and initialize a new instance of Config from env vars
func NewConfig() (*Config, error) {

	var (
		err error
		cfg Config
	)
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("envconfig.Process() failed: %w", err)
	}

	for k, v := range cfg.HttpEndpointTopic {
		cfg.HttpEndpointTopic[k] = strings.ReplaceAll(v, "_", ":")
	}

	return &cfg, nil
}
