package http

import (
	"io/ioutil"
	"log"
	"net/http"
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
