package vk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// handle all subscription confirmations
func (hh *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	decoded := SubscriptionConfirmation{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		log.Printf("%s %s json.Decode() failed: %v", ident, r.RemoteAddr, err)
		hh.httpError(w, http.StatusBadRequest,
			"Failed to decode JSON in the body: %s", err)
		return
	}

	reqProto := r.Header.Get("X-Forwarded-Proto")
	if reqProto == "" {
		reqProto = "http"
	}

	fullURL := fmt.Sprintf("%s://%s%s", reqProto, r.Host, r.URL.String())
	if hh.debug {
		log.Printf("%s %s subscription confirmation request for: %s\n",
			ident, r.RemoteAddr, fullURL)
	}

	response := struct {
		Signature string `json:"signature"` // base64-encoded signature
	}{
		Signature: signSubscriptionHex(decoded, fullURL),
	}
	w.Header().Set("Content-Type", "application/json")
	jsonBody, err := json.Marshal(&response)
	if err != nil {
		log.Printf("%s %s json.Marshal() failed: %v", ident, r.RemoteAddr, err)
		hh.httpError(w, http.StatusInternalServerError,
			"Failed to marshal json: %s", err)
		return
	}

	if hh.debug {
		log.Printf("%s %s subscription confirmation response for: %s\n%s\n",
			ident, r.RemoteAddr, fullURL, string(jsonBody))
	}

	w.Write(jsonBody)
}

//
func (hh *Handler) httpError(
	w http.ResponseWriter,
	code int, format string, args ...interface{}) {

	http.Error(w, fmt.Sprintf(format, args...), code)
}
