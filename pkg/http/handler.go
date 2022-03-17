package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// handler passes notification to SNS.
//  URL is a SNS topic.
func (s *srv) handler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s %s read body failed: %v\n", ident, r.RemoteAddr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		log.Printf("%s %s empty body", ident, r.RemoteAddr)
	}

	if topic, ok := s.endpoint2topic[r.URL.Path]; ok {
		if err = s.sns.Publish(topic, string(data)); err != nil {
			log.Printf("%s %s Publish() failed: %v", ident, r.RemoteAddr, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("%s %s publish event to %s\n", ident, r.RemoteAddr, topic)
		return
	}

	log.Printf("%s %s bad request %s\n", ident, r.RemoteAddr, r.URL.Path)

	// slow bad request reply
	time.Sleep(time.Second)
	http.Error(w, "bad request", http.StatusBadRequest)
}

// Middleware to handle subscription confirmation
func (s *srv) subscriptionConfimation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s.dump(r)

		// X-Amz-Sns-Message-Type: SubscriptionConfirmation
		snsType := strings.ToLower(r.Header.Get("X-Amz-Sns-Message-Type"))
		switch snsType {
		case "subscriptionconfirmation":
			s.sch.Handle(w, r, s.debug)
		default:
			next.ServeHTTP(w, r)
		}
	})
}

//
func (s *srv) dump(r *http.Request) {
	if !s.debug {
		return
	}

	// dump url
	reqProto := r.Header.Get("X-Forwarded-Proto")
	if reqProto == "" {
		reqProto = "http"
	}
	fullURL := fmt.Sprintf("%s://%s%s", reqProto, r.Host, r.URL.String())
	log.Printf("%s %s got http request %s\n", ident, r.RemoteAddr, fullURL)

	// dump request
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("%s %s failed to dump requst: %v\n", ident, r.RemoteAddr, err)
		return
	}
	log.Printf("%s %s\n%s\n", ident, r.RemoteAddr, dump)
}
