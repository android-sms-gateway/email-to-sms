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
  <a href="https://github.com/capcom6/go-project-template">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" alt="Logo" width="120" height="120">
  </a>

<h3 align="center">go-project-template</h3>

  <p align="center">
    Opinionated Go service template with Fiber API, OpenAPI docs, Telegram bot wiring, health endpoints, and Fx-based modular DI.
    <br />
    <a href="https://github.com/capcom6/go-project-template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/capcom6/go-project-template/issues">Report Bug</a>
    ·
    <a href="https://github.com/capcom6/go-project-template/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)



<!-- ABOUT THE PROJECT -->
## About The Project

This repository is a production-oriented starter for backend services in Go. It ships with:

* HTTP server bootstrapped with Fiber and dependency injection via Uber Fx.
* Version + health endpoints using `healthfx`.
* Swagger/OpenAPI docs endpoint under `/api/v1/docs`.
* Telegram bot integration (`/start` command handler included).
* Modular business domain example to plug in real use-cases.

Use this template when you want a fast path to shipping APIs and bot workflows with a clean module layout.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* [![Go][go-shield]][go-url]
* [![Fiber][fiber-shield]][fiber-url]
* [![Fx][fx-shield]][fx-url]
* [![Swagger][swagger-shield]][swagger-url]
* [![Telegram][telegram-shield]][telegram-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Follow these steps to run the service locally.

### Prerequisites

* Go 1.25+
  ```sh
  go version
  ```
* `golangci-lint` (optional but recommended)
  ```sh
  golangci-lint version
  ```
* `swag` CLI for docs generation
  ```sh
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

### Installation

1. Clone the repo.
   ```sh
   git clone https://github.com/capcom6/go-project-template.git
   cd go-project-template
   ```
2. Download dependencies.
   ```sh
   make deps
   ```
3. (Optional) Generate OpenAPI docs.
   ```sh
   make gen
   ```
4. Build the binary.
   ```sh
   make build
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Run the app:

```sh
go run .
```

Or use live reload:

```sh
make air
```

Default server address is `127.0.0.1:3000`.

Helpful endpoints:

* Health endpoints (via `healthfx`), typically under `/health`.
* OpenAPI docs: `http://127.0.0.1:3000/api/v1/docs`.

Configuration:

* Environment variables are loaded via `go-core-fx/config`.
* You can point to a YAML file using:
  ```sh
  export CONFIG_PATH=./config.local.yaml
  ```
* Telegram token is configured via the app config (`telegram.token`) and required to use bot handlers.

Quality checks:

```sh
make fmt
make lint
make test
make coverage
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Fiber API scaffold
- [x] OpenAPI + Swagger integration
- [x] Telegram bot wiring
- [ ] Database module integration
- [ ] Auth middleware and RBAC
- [ ] CI release pipeline hardening

See the [open issues](https://github.com/capcom6/go-project-template/issues) for a full list of proposed features (and known issues).

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

Project Link: [https://github.com/capcom6/go-project-template](https://github.com/capcom6/go-project-template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Go Fiber](https://github.com/gofiber/fiber)
* [Uber Fx](https://github.com/uber-go/fx)
* [Telego](https://github.com/mymmrac/telego)
* [Shields.io](https://shields.io)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/capcom6/go-project-template.svg?style=for-the-badge
[contributors-url]: https://github.com/capcom6/go-project-template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/capcom6/go-project-template.svg?style=for-the-badge
[forks-url]: https://github.com/capcom6/go-project-template/network/members
[stars-shield]: https://img.shields.io/github/stars/capcom6/go-project-template.svg?style=for-the-badge
[stars-url]: https://github.com/capcom6/go-project-template/stargazers
[issues-shield]: https://img.shields.io/github/issues/capcom6/go-project-template.svg?style=for-the-badge
[issues-url]: https://github.com/capcom6/go-project-template/issues
[license-shield]: https://img.shields.io/github/license/capcom6/go-project-template.svg?style=for-the-badge
[license-url]: https://github.com/capcom6/go-project-template/blob/master/LICENSE
[go-shield]: https://img.shields.io/badge/go-1.25%2B-00ADD8?style=for-the-badge&logo=go
[go-url]: https://go.dev/
[fiber-shield]: https://img.shields.io/badge/Fiber-v2-00b894?style=for-the-badge
[fiber-url]: https://github.com/gofiber/fiber
[fx-shield]: https://img.shields.io/badge/Uber%20Fx-DI-6f42c1?style=for-the-badge
[fx-url]: https://github.com/uber-go/fx
[swagger-shield]: https://img.shields.io/badge/OpenAPI-Swagger-85EA2D?style=for-the-badge
[swagger-url]: https://github.com/swaggo/swag
[telegram-shield]: https://img.shields.io/badge/Telegram-Bot-26A5E4?style=for-the-badge&logo=telegram
[telegram-url]: https://core.telegram.org/bots
