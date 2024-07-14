# wgezpeer

## Overview

`wgezpeer` is a Go-based utility for generating WireGuard peer configurations. It takes a CIDR and an endpoint as input arguments and outputs the server and client configuration files needed to establish a WireGuard VPN connection.

## Features

- Generates WireGuard server and client configuration files.
- Automatically increments IP addresses within the given CIDR.
- Randomly selects a listening port for the client.
- Uses the WireGuard Go library for key generation and configuration.

## Requirements

- Go 1.22.4 or later
- WireGuard Go library

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/alirezasn3/wgezpeer.git
   cd wgezpeer
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

To generate WireGuard configurations, run the following command:

```sh
go run main.go <CIDR> <Endpoint>
```

Example:

```sh
go run main.go 10.0.0.1/24 12.23.34.45
```

This will output the server and client configuration files.

## Files

### `main.go`

The main entry point of the application. It handles the following:

- Parsing command-line arguments.
- Validating the CIDR.
- Generating server and client IP addresses.
- Creating WireGuard configurations.
- Printing the configurations to the console.

### `IPAddress.go`

Contains the `IPAddress` struct and methods for:

- Parsing an IP address from a string.
- Incrementing the IP address.
- Converting the IP address to a string.

### `go.mod`

Defines the module and its dependencies.

### `go.sum`

Contains checksums for module dependencies.

## Dependencies

- `golang.zx2c4.com/wireguard/wgctrl`
- `github.com/google/go-cmp`
- `github.com/josharian/native`
- `github.com/mdlayher/genetlink`
- `github.com/mdlayher/netlink`
- `github.com/mdlayher/socket`
- `golang.org/x/crypto`
- `golang.org/x/net`
- `golang.org/x/sync`
- `golang.org/x/sys`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Contact

For any questions or issues, please open an issue on the [GitHub repository](https://github.com/alirezasn3/wgezpeer).

---

This README provides a comprehensive overview of the `wgezpeer` project, including its purpose, usage, and structure. Feel free to modify it to better suit your needs.
