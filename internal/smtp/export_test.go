package smtp

import gosmtp "github.com/emersion/go-smtp"

// ServerForTest exposes the underlying go-smtp server so tests can serve it on
// a controlled listener and drive it over the wire.
func ServerForTest(s *Server) *gosmtp.Server {
	return s.server
}
