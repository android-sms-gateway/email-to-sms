package smtp

import "context"

// nopHandler is a temporary no-op handler used until the bridge module is implemented.
type nopHandler struct{}

func (nopHandler) Handle(_ context.Context, _ EmailData) error {
	return nil
}

// NewNopHandler returns a temporary no-op Handler.
func NewNopHandler() Handler {
	return nopHandler{}
}
