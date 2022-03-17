package vk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// handle all subscription confirmations
func (hh *Handler) Handle(w http.ResponseWriter, r *http.Request, debug bool) {

	defer r.Body.Close()

	decoded := SubscriptionConfirmation{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		hh.httpError(w, debug, http.StatusBadRequest,
			"Failed to decode JSON in the body: %s", err)
		return
	}

	reqProto := r.Header.Get("X-Forwarded-Proto")
	if reqProto == "" {
		reqProto = "http"
	}

	fullURL := fmt.Sprintf("%s://%s%s", reqProto, r.Host, r.URL.String())
	if debug {
		log.Printf("%s subscription confirmation request for: %s\n",
			ident, fullURL)
	}

	response := struct {
		Signature string `json:"signature"` // base64-encoded signature
	}{
		Signature: signSubscriptionHex(decoded, fullURL),
	}
	w.Header().Set("Content-Type", "application/json")
	jsonBody, err := json.Marshal(&response)
	if err != nil {
		hh.httpError(w, debug, http.StatusInternalServerError,
			"Failed to marshal json: %s", err)
		return
	}

	if debug {
		log.Printf("%s subscription configmation response for: %s\n%s\n",
			ident, fullURL, string(jsonBody))
	}

	w.Write(jsonBody)
}

//
func (hh *Handler) httpError(
	w http.ResponseWriter,
	debug bool,
	code int, format string, args ...interface{}) {

	text := fmt.Sprintf(format, args...)
	if debug {
		log.Printf("%s %s\n", ident, text)
	}
	http.Error(w, text, code)
}
