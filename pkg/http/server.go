package http

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/tdx/http2sns-vk/pkg/config"
	snsApi "github.com/tdx/http2sns-vk/pkg/sns/api"
	subApi "github.com/tdx/http2sns-vk/pkg/subscription/api"
)

const ident = "[HTTP]:"

type srv struct {
	debug          bool
	endpoint2topic map[string]string
	mux            *http.ServeMux
	sch            subApi.SubscriptionConfirmationHandler
	sns            snsApi.Publisher
}

// Start runs HTTP server with specified endpoints and
//  subscription confirmation handler
func Start(
	cfg *config.Config,
	sch subApi.SubscriptionConfirmationHandler,
	sns snsApi.Publisher) {

	l, err := net.Listen("tcp4", cfg.HttpListenAddr)
	if err != nil {
		log.Fatalf("%s net.Listen(%s) failed: %v\n",
			ident, cfg.HttpListenAddr, err)
	}

	s := &srv{
		debug:          cfg.HttpDebug,
		endpoint2topic: cfg.HttpEndpointTopic,
		mux:            http.NewServeMux(),
		sch:            sch,
		sns:            sns,
	}

	s.initEndpoints()

	log.Printf("%s start HTTP server on %s", ident, cfg.HttpListenAddr)
	log.Fatal(http.Serve(l, s.mux))
}

//
func (s *srv) initEndpoints() {
	for endpoint, topic := range s.endpoint2topic {
		if !strings.HasPrefix(endpoint, "/") {
			endpoint = "/" + endpoint
		}
		if s.debug {
			log.Printf("%s registering '%s' -> '%s'\n",
				ident, endpoint, topic)
		}
		s.mux.Handle(endpoint,
			DumpRequest(s.debug,
				SubscriptionConfimaton(s.sch,
					http.HandlerFunc(s.handler)),
			),
		)
	}
}
