package api

import "net/http"

// SubscriptionConfirmationHandler specifies interface for handler used by
//  http server to confirm subscription
type SubscriptionConfirmationHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
