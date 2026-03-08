
# 🔒 GopherTLS

A wrapper for [tls-client](https://github.com/bogdanfinn/tls-client) library.

## 📝 Description

An API that forwards your HTTP requests using a custom TLS fingerprint.

## 🚀 Installation

1. `git clone https://github.com/RealSadmc/GopherTLS-API.git`
2. `cd GopherTLS-API`
3. `go run .`

> [!TIP]
> Configure the API server host and port through the `.env` file.

## 📚 Documentation

### Endpoint: `/go/forward`

### Method: `POST`

### Headers:

| Header                              | Description                                                                                                                                                            | Optional | Default |
|-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|---------|
| `x-tls-url`                         | 🌐 Request URL.                                                                                                                                                        | No       | `N/A`   |
| `x-tls-method`                      | 📮 Request method.                                                                                                                                                     | No       | `N/A`   |
| `x-tls-proxy`                       | 🔄 Proxy. Formats: `ip:port:user:pass`, `ip:port`                                                                                                                      | Yes      | `N/A`   |
| `x-tls-profile`                     | 👤 TLS client profile. Available profiles: [See here](https://raw.githubusercontent.com/bogdanfinn/tls-client/refs/heads/master/profiles/profiles.go) | No       | `N/A`   |
| `x-tls-timeout`              | ⏱️ HTTP client timeout.                                                                                                                                                | Yes      | `30`    |
| `x-tls-follow-redirects`            | 🔀 Follow redirects.                                                                                                                                                   | Yes      | `true`  |
| `x-tls-force-http1`                 | 🔌 Force HTTP/1.1                                                                                                                                                           | Yes      | `false` |
| `x-tls-insecure-skip-verify`        | 🚫 Skip SSL certificate verification.                                                                                                                                  | Yes      | `false` |
| `x-tls-with-random-extension-order` | 🎲 Randomize extensions order.                                                                                                                                         | Yes      | `true`  |
| `x-tls-header-order`                | 📋 Header order. Format: String with headers key separated by commas (`,`)                                                                                             | Yes      | `N/A`   |
| `x-tls-pseudo-header-order`         | 📑 Pseudo header order. Format: String with headers key separated by commas (`,`)                                                                                      | Yes      | `N/A`   |

> [!NOTE]
> If the request requires a body, you can simply enter it as the request body, not in the header.

## 🐛 Report Issues

Found a bug? Please [Report The issue](https://github.com/RealSadmc/GopherTLS-API/issues).


By reporting an issue you help improve the project.

## 🙏 Credits

This is A Modified Fork Of [brianxor](https://github.com/brianxor/tls-api) Shoutout to you my man

Special thanks to [bogdanfinn](https://github.com/bogdanfinn/) for creating the awesome [tls-client](https://github.com/bogdanfinn/tls-client)
