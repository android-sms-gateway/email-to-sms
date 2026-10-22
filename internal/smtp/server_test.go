package smtp_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	netsmtp "net/smtp"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/android-sms-gateway/email-to-sms/internal/smtp"
)

func TestServerTLSConfiguration(t *testing.T) {
	t.Parallel()

	metrics := smtp.NewMetrics()
	handler := smtp.NewNopHandler()
	logger := zap.NewNop()

	t.Run("without TLS", func(t *testing.T) {
		t.Parallel()

		server, err := smtp.NewServer(smtp.Config{
			Host:    "127.0.0.1",
			Port:    0,
			Domain:  "example.com",
			TLSCert: "",
			TLSKey:  "",
		}, handler, metrics, logger)
		if err != nil {
			t.Fatalf("NewServer: %v", err)
		}

		addr, shutdown := startServer(t, server)
		defer shutdown()

		c, err := netsmtp.Dial(addr)
		if err != nil {
			t.Fatalf("dial: %v", err)
		}
		defer c.Close()

		err = c.Hello("localhost")
		if err != nil {
			t.Fatalf("hello: %v", err)
		}

		if ok, _ := c.Extension("AUTH"); !ok {
			t.Fatal("expected AUTH to be advertised without TLS configured")
		}
		if ok, _ := c.Extension("STARTTLS"); ok {
			t.Fatal("did not expect STARTTLS without TLS configured")
		}

		err = c.Auth(netsmtp.PlainAuth("", "user", "pass", "127.0.0.1"))
		if err != nil {
			t.Fatalf("cleartext auth should succeed without TLS configured: %v", err)
		}
	})

	t.Run("with TLS", func(t *testing.T) {
		t.Parallel()

		certPath, keyPath := writeSelfSignedCert(t)
		server, err := smtp.NewServer(smtp.Config{
			Host:    "127.0.0.1",
			Port:    0,
			Domain:  "example.com",
			TLSCert: certPath,
			TLSKey:  keyPath,
		}, handler, metrics, logger)
		if err != nil {
			t.Fatalf("NewServer: %v", err)
		}

		addr, shutdown := startServer(t, server)
		defer shutdown()

		assertCleartextAuthRefused(t, addr)

		pool := x509.NewCertPool()
		certPEM, err := os.ReadFile(certPath)
		if err != nil {
			t.Fatalf("read cert: %v", err)
		}
		if !pool.AppendCertsFromPEM(certPEM) {
			t.Fatal("failed to append test certificate to pool")
		}

		c, err := netsmtp.Dial(addr)
		if err != nil {
			t.Fatalf("dial: %v", err)
		}
		defer c.Close()

		err = c.Hello("localhost")
		if err != nil {
			t.Fatalf("hello: %v", err)
		}
		err = c.StartTLS(&tls.Config{RootCAs: pool, ServerName: "127.0.0.1"})
		if err != nil {
			t.Fatalf("starttls: %v", err)
		}
		if ok, _ := c.Extension("AUTH"); !ok {
			t.Fatal("expected AUTH to be advertised after STARTTLS")
		}

		err = c.Auth(netsmtp.PlainAuth("", "user", "pass", "127.0.0.1"))
		if err != nil {
			t.Fatalf("auth after STARTTLS should succeed: %v", err)
		}
	})

	t.Run("incomplete TLS configuration", func(t *testing.T) {
		t.Parallel()

		certPath, keyPath := writeSelfSignedCert(t)

		cases := []struct {
			name    string
			tlsCert string
			tlsKey  string
		}{
			{name: "certificate without key", tlsCert: certPath},
			{name: "key without certificate", tlsKey: keyPath},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				server, err := smtp.NewServer(smtp.Config{
					Host:    "127.0.0.1",
					Port:    0,
					Domain:  "example.com",
					TLSCert: tc.tlsCert,
					TLSKey:  tc.tlsKey,
				}, handler, metrics, logger)
				if !errors.Is(err, smtp.ErrIncompleteTLSConfig) {
					t.Fatalf("expected ErrIncompleteTLSConfig, got %v", err)
				}
				if server != nil {
					t.Fatal("expected no server for incomplete TLS configuration")
				}
			})
		}
	})
}

// assertCleartextAuthRefused connects without TLS and verifies the server does
// not advertise AUTH and rejects a forced AUTH PLAIN command.
func assertCleartextAuthRefused(t *testing.T, addr string) {
	t.Helper()

	c, err := netsmtp.Dial(addr)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	err = c.Hello("localhost")
	if err != nil {
		t.Fatalf("hello: %v", err)
	}
	if ok, _ := c.Extension("STARTTLS"); !ok {
		t.Fatal("expected STARTTLS to be advertised with TLS configured")
	}
	if ok, _ := c.Extension("AUTH"); ok {
		t.Fatal("AUTH must not be advertised over cleartext when TLS is configured")
	}

	err = c.Auth(netsmtp.PlainAuth("", "user", "pass", "127.0.0.1"))
	if err == nil {
		t.Fatal("expected cleartext AUTH to be refused when TLS is configured")
	}
}

func startServer(t *testing.T, server *smtp.Server) (string, func()) {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	go func() {
		_ = smtp.ServerForTest(server).Serve(listener)
	}()

	return listener.Addr().String(), func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_ = smtp.ServerForTest(server).Shutdown(ctx)
	}
}

func writeSelfSignedCert(t *testing.T) (string, string) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}

	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}

	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("create certificate: %v", err)
	}

	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatalf("marshal key: %v", err)
	}

	dir := t.TempDir()
	certPath := filepath.Join(dir, "cert.pem")
	keyPath := filepath.Join(dir, "key.pem")

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})

	err = os.WriteFile(certPath, certPEM, 0o600)
	if err != nil {
		t.Fatalf("write certificate: %v", err)
	}
	err = os.WriteFile(keyPath, keyPEM, 0o600)
	if err != nil {
		t.Fatalf("write key: %v", err)
	}

	return certPath, keyPath
}
