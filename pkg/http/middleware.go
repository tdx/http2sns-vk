package http

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	subApi "github.com/tdx/http2sns-vk/pkg/subscription/api"
)

// SubscriptionConfimation is a HTTP middleware to handle subscription confirmation
func SubscriptionConfirmaton(
	sch subApi.SubscriptionConfirmationHandler,
	next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// X-Amz-Sns-Message-Type: SubscriptionConfirmation
		snsType := strings.ToLower(r.Header.Get("X-Amz-Sns-Message-Type"))
		switch snsType {
		case "subscriptionconfirmation":
			sch.Handle(w, r)
		default:
			next.ServeHTTP(w, r)
		}
	})
}

// DumpRequest is a HTTP middleware to dump HTTP request
func DumpRequest(dump bool, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if dump {
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

		next.ServeHTTP(w, r)
	})
}
