package vk

// Handler ..
type Handler struct {
	debug bool
}

const ident = "[VK]:"

// NewHandler returns SubscriptionConfirmationHandler
//  implementation for VK Cloud
func NewHandler(debug bool) *Handler {
	return &Handler{debug}
}
