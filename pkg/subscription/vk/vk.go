package vk

// Handler ..
type Handler struct {
}

const ident = "[VK]:"

// NewHandler returns SubscriptionConfirmationHandler
//  implementation for VK Cloud
func NewHandler() *Handler {
	return &Handler{}
}
