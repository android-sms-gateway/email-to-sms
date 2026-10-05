# 📱 Email-to-SMS Bridge

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stars][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![License][license-shield]][license-url]

An SMTP server that receives emails and forwards them as SMS messages through the SMSGate REST API. See [Email to SMS](https://docs.sms-gate.app/services/email-to-sms/) for the full guide.

## 📖 About

The Email-to-SMS Bridge lets any existing email-based notification system send SMS with zero code changes: applications send email to `phone@smtp-domain`, and the bridge forwards the message body to the SMSGate API using the SMTP AUTH credentials passed through for authentication.

## 📚 Table of Contents

- [📖 About](#-about)
- [📚 Table of Contents](#-table-of-contents)
- [⭐ Features](#-features)
- [📦 Prerequisites](#-prerequisites)
- [🚀 Getting Started](#-getting-started)
- [⚙️ Configuration](#️-configuration)
- [📦 Deployment](#-deployment)
- [🔌 API Overview](#-api-overview)
- [📚 Documentation](#-documentation)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

## ⭐ Features

- SMTP server with `AUTH PLAIN` and optional STARTTLS
- Recipient-domain validation and phone number extraction from the local part
- Credentials passed through to SMSGate for per-user authentication
- Prometheus metrics (emails received, SMS sent/failed, auth failures)
- HTTP API with Swagger/OpenAPI docs

## 📦 Prerequisites

- Go 1.25+ for building from source
- An SMSGate account
- SMTP client that supports `AUTH PLAIN` (most email systems)

## 🚀 Getting Started

```bash
make build
SMTP__DOMAIN=sms-gateway.local ./bin/email-to-sms
```

Send an email to `79991234567@sms-gateway.local`; the message body becomes the SMS text.

## ⚙️ Configuration

Configuration is loaded from environment variables (double-underscore nesting) or an optional YAML file via `CONFIG_PATH`.

| Variable                     | Default                                | Description                                   |
| ---------------------------- | -------------------------------------- | --------------------------------------------- |
| `HTTP__ADDRESS`              | `127.0.0.1:3000`                       | HTTP API listen address                       |
| `HTTP__OPENAPI__ENABLED`     | `true`                                 | Enable Swagger UI at `/api/v1/docs`           |
| `SMTP__HOST`                 | `127.0.0.1`                            | SMTP listen address                           |
| `SMTP__PORT`                 | `587`                                  | SMTP port                                     |
| `SMTP__DOMAIN`               | `example.com`                          | Allowed recipient domain                      |
| `SMTP__TLS_CERT`             | *(empty)*                              | TLS cert path (STARTTLS)                      |
| `SMTP__TLS_KEY`              | *(empty)*                              | TLS key path                                  |
| `SMSGATE__URL`               | `https://api.sms-gate.app/3rdparty/v1` | SMSGate API base URL                          |
| `SMSGATE__SKIP_PHONE_VALIDATION` | `false`                            | Skip phone number validation                  |
| `CONFIG_PATH`                | *(empty)*                              | Optional YAML config file path                |

Full reference: [Email to SMS](https://docs.sms-gate.app/services/email-to-sms/).

## 📦 Deployment

```bash
docker run -d --name email-to-sms \
  -e SMTP__DOMAIN=sms-gateway.local \
  -e SMTP__HOST=0.0.0.0 \
  -e SMTP__PORT=2525 \
  -e HTTP__ADDRESS=0.0.0.0:3000 \
  -p 2525:2525 \
  -p 3000:3000 \
  ghcr.io/android-sms-gateway/email-to-sms:latest
```

> **Note:** SMTP `AUTH PLAIN` is accepted before STARTTLS, so credentials cross the published port in cleartext. On untrusted networks, restrict the published port (e.g. `-p 127.0.0.1:2525:2525`) or terminate TLS in front of the container.

## 🔌 API Overview

| Method | Path              | Description                                   |
| ------ | ----------------- | --------------------------------------------- |
| GET    | `/health`         | Liveness probe (same as `/health/live`)       |
| GET    | `/health/live`    | Liveness probe                                |
| GET    | `/health/ready`   | Readiness probe                               |
| GET    | `/health/startup` | Startup probe                                 |
| GET    | `/metrics`        | Prometheus metrics                            |
| GET    | `/api/v1/docs`    | Swagger/OpenAPI documentation                 |

## 📚 Documentation

- [Email to SMS](https://docs.sms-gate.app/services/email-to-sms/)
- [SMTP integration](https://docs.sms-gate.app/integration/smtp/)
- [Central docs](https://docs.sms-gate.app/)

## 🤝 Contributing

Contributions are welcome. Open an issue or pull request; follow the repo's `make fmt` / `make lint` / `make test` checks and the [README style guide](https://docs.sms-gate.app/).

## 📄 License

Apache-2.0. See [LICENSE](LICENSE).

<!-- Reference-style badge URLs: style=for-the-badge is mandatory -->
[contributors-shield]: https://img.shields.io/github/contributors/android-sms-gateway/email-to-sms?style=for-the-badge
[contributors-url]: https://github.com/android-sms-gateway/email-to-sms/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/android-sms-gateway/email-to-sms?style=for-the-badge
[forks-url]: https://github.com/android-sms-gateway/email-to-sms/network/members
[stars-shield]: https://img.shields.io/github/stars/android-sms-gateway/email-to-sms?style=for-the-badge
[stars-url]: https://github.com/android-sms-gateway/email-to-sms/stargazers
[issues-shield]: https://img.shields.io/github/issues/android-sms-gateway/email-to-sms?style=for-the-badge
[issues-url]: https://github.com/android-sms-gateway/email-to-sms/issues
[license-shield]: https://img.shields.io/github/license/android-sms-gateway/email-to-sms?style=for-the-badge
[license-url]: https://github.com/android-sms-gateway/email-to-sms/blob/master/LICENSE
