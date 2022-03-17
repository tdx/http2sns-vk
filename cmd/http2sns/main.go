package main

import (
	"log"
	"os"

	"github.com/tdx/http2sns-vk/pkg/config"
	"github.com/tdx/http2sns-vk/pkg/http"
	"github.com/tdx/http2sns-vk/pkg/sns/aws"
	"github.com/tdx/http2sns-vk/pkg/subscription/vk"
)

const ident = "[MAIN]:"

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("%s %v\n", ident, err)
		os.Exit(1)
	}
	dumpConfig(cfg)

	snsPublisher, err := aws.NewPublisher(cfg)
	if err != nil {
		log.Printf("%s %v\n", ident, err)
		os.Exit(1)
	}
	subHandler := vk.NewHandler(cfg.HttpDebug)
	http.Start(cfg, subHandler, snsPublisher)
}

func dumpConfig(cfg *config.Config) {
	log.Printf("%s http.listen.addr: %s\n", ident, cfg.HttpListenAddr)
	log.Printf("%s http.endpoint.topic: %v\n", ident, cfg.HttpEndpointTopic)
	log.Printf("%s http.debug: %v\n", ident, cfg.HttpDebug)
	log.Printf("%s sns.region: %v\n", ident, cfg.SnsRegion)
	log.Printf("%s sns.api.endpoint: %v\n", ident, cfg.SnsApiEndpoint)
}
