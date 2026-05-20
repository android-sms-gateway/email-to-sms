package smtp

// Config holds configuration for the SMTP server module.
type Config struct {
	Host    string
	Port    int
	Domain  string
	TLSCert string
	TLSKey  string
}
