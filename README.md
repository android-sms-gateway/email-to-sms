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
    A standalone SMTP server that receives emails and forwards them as SMS messages via the SMSGate platform.
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
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)



<!-- ABOUT THE PROJECT -->
## About The Project

The Email-to-SMS Bridge is a standalone SMTP server that receives emails and forwards them as SMS messages via the SMSGate Android gateway.

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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

WIP

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

WIP

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Project initialization
- [ ] SMTP server implementation
- [ ] Email parsing
- [ ] SMS sending integration
- [ ] Docker support

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

* [SMSGate Platform](https://sms-gate.app)
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
[go-shield]: https://img.shields.io/badge/go-1.21%2B-00ADD8?style=for-the-badge&logo=go
[go-url]: https://go.dev/
