<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache License][license-shield]][license-url]
[![Go][go-shield]][go-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/android-sms-gateway/email-to-sms">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" alt="Logo" width="120" height="120">
  </a>

  <h3 align="center">Email-to-SMS Bridge</h3>

  <p align="center">
    SMTP server that receives emails and forwards them as SMS messages via the SMSGate Android gateway.
    <br />
    <a href="https://github.com/android-sms-gateway/email-to-sms"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/android-sms-gateway/email-to-sms/issues">Report Bug</a>
    ·
    <a href="https://github.com/android-sms-gateway/email-to-sms/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
- [About The Project](#about-the-project)
  - [Use Case](#use-case)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Pre-built binaries (GitHub Releases)](#pre-built-binaries-github-releases)
  - [Docker (GHCR)](#docker-ghcr)
  - [Building from source](#building-from-source)
  - [Development](#development)
- [Usage](#usage)
  - [Sending SMS via email](#sending-sms-via-email)
  - [Configuration](#configuration)
  - [SMTP Authentication](#smtp-authentication)
  - [Monitoring](#monitoring)
  - [Docker](#docker)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)



<!-- ABOUT THE PROJECT -->
## About The Project

The Email-to-SMS Bridge is a standalone SMTP server that receives emails and forwards them as SMS messages via the [SMSGate](https://sms-gate.app) Android gateway. Built with Go, it uses [go-smtp](https://github.com/emersion/go-smtp) for SMTP, [Fiber](https://gofiber.io/) for HTTP API, [Prometheus](https://prometheus.io/) for metrics, and [uber-go/fx](https://uber-go.github.io/fx/) for dependency injection.

This service acts as a bridge between email-based notification systems and SMS delivery, allowing booking systems, appointment schedulers, and other applications to send SMS notifications through a simple email interface.

### Use Case

Applications that already send email notifications can easily add SMS delivery by sending emails to this bridge:

```
Booking System → Email → Email-to-SMS Bridge → SMSGate → Android Device → SMS
```

Emails are sent in the format: `79991234567@sms-gateway.local`

- **Local part (before @):** Phone number in any format
- **Body:** SMS message content
- **Authentication:** SMTP AUTH credentials are passed through to SMSGate for per-user authentication

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* [![Go][go-shield]][go-url]
* [go-smtp](https://github.com/emersion/go-smtp) — SMTP server library
* [SMSGate client-go](https://github.com/android-sms-gateway/client-go) — SMS gateway API client
* [Fiber](https://gofiber.io/) — HTTP framework
* [uber-go/fx](https://uber-go.github.io/fx/) — Dependency injection
* [Prometheus](https://prometheus.io/) — Metrics and monitoring
* [go-core-fx/config](https://github.com/go-core-fx/config) — Configuration management

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Pre-built binaries (GitHub Releases)

Download the latest binary for your platform from the [Releases page](https://github.com/android-sms-gateway/email-to-sms/releases).

```sh
# Example: Linux amd64
curl -LO https://github.com/android-sms-gateway/email-to-sms/releases/latest/download/email-to-sms_linux_amd64.tar.gz
tar xzf email-to-sms_linux_amd64.tar.gz

# Run
SMTP__DOMAIN=sms-gateway.local ./email-to-sms
```

### Docker (GHCR)

```sh
docker pull ghcr.io/android-sms-gateway/email-to-sms:latest

docker run -d \
  -e SMTP__DOMAIN=sms-gateway.local \
  -e SMTP__PORT=2525 \
  -p 2525:2525 \
  ghcr.io/android-sms-gateway/email-to-sms:latest
```

### Building from source

```sh
git clone https://github.com/android-sms-gateway/email-to-sms.git
cd email-to-sms
make build
```

Requires Go 1.25+ (see `go.mod`). The binary will be placed in `bin/email-to-sms`.

### Development

```sh
# Install Air for live reload
go install github.com/air-verse/air@latest

# Start dev server with hot reload
make air
```

See the [Makefile](./Makefile) for all available targets (`make help`).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

### Sending SMS via email

Send an email where the recipient address contains the phone number:

```
To: 79991234567@sms-gateway.local
Subject: (ignored)
Body: Your appointment is confirmed for tomorrow at 10:00 AM.
```

### Configuration

The application is configured via environment variables (double-underscore for nesting) or a YAML file specified by `CONFIG_PATH`.

| Variable                         | Default                                | Description                         |
| -------------------------------- | -------------------------------------- | ----------------------------------- |
| `HTTP__ADDRESS`                  | `127.0.0.1:3000`                       | HTTP API listen address             |
| `HTTP__OPENAPI__ENABLED`         | `true`                                 | Enable Swagger UI at `/api/v1/docs` |
| `SMTP__HOST`                     | `127.0.0.1`                            | SMTP listen address                 |
| `SMTP__PORT`                     | `587`                                  | SMTP port                           |
| `SMTP__DOMAIN`                   | `example.com`                          | Allowed domain for email recipients |
| `SMTP__TLS_CERT`                 | `""`                                   | Path to TLS certificate (optional)  |
| `SMTP__TLS_KEY`                  | `""`                                   | Path to TLS key (optional)          |
| `SMSGATE__URL`                   | `https://api.sms-gate.app/3rdparty/v1` | SMSGate API base URL                |
| `SMSGATE__SKIP_PHONE_VALIDATION` | `false`                                | Skip phone validation on SMSGate    |

### SMTP Authentication

SMTP AUTH PLAIN credentials are passed through to SMSGate for per-user authentication.

### Monitoring

Prometheus metrics are exposed on the HTTP server at `/metrics`:

- `email2sms_bridge_emails_received_total` — Emails received by the SMTP server
- `email2sms_smsgate_sms_sent_total` — Successful SMS sends
- `email2sms_smsgate_sms_failed_total` — Failed SMS sends
- `email2sms_smsgate_auth_failures_total` — Authentication failures

### Docker

```sh
make docker-build
```

A production Docker image is also built via GoReleaser on every release.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Project initialization
- [x] SMTP server implementation (AUTH PLAIN, TLS, domain validation)
- [x] Email parsing (MIME, phone extraction)
- [x] SMS sending via SMSGate API
- [x] Prometheus metrics
- [x] HTTP API with Swagger docs
- [x] Docker support (multi-arch)
- [ ] Unit and integration tests
- [ ] Helm chart for Kubernetes deployment

See the [open issues](https://github.com/android-sms-gateway/email-to-sms/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Maintainer: [@capcom6](https://github.com/capcom6)

Project Link: [https://github.com/android-sms-gateway/email-to-sms](https://github.com/android-sms-gateway/email-to-sms)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Go](https://go.dev/)
* [Shields.io](https://shields.io)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/android-sms-gateway/email-to-sms.svg?style=for-the-badge
[contributors-url]: https://github.com/android-sms-gateway/email-to-sms/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/android-sms-gateway/email-to-sms.svg?style=for-the-badge
[forks-url]: https://github.com/android-sms-gateway/email-to-sms/network/members
[stars-shield]: https://img.shields.io/github/stars/android-sms-gateway/email-to-sms.svg?style=for-the-badge
[stars-url]: https://github.com/android-sms-gateway/email-to-sms/stargazers
[issues-shield]: https://img.shields.io/github/issues/android-sms-gateway/email-to-sms.svg?style=for-the-badge
[issues-url]: https://github.com/android-sms-gateway/email-to-sms/issues
[license-shield]: https://img.shields.io/github/license/android-sms-gateway/email-to-sms.svg?style=for-the-badge
[license-url]: https://github.com/android-sms-gateway/email-to-sms/blob/master/LICENSE
[go-shield]: https://img.shields.io/badge/go-1.25%2B-00ADD8?style=for-the-badge&logo=go
[go-url]: https://go.dev/
