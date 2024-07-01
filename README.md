[![Build status][status-shield]][status-url]
[![Linter status][linter-status-shield]][linter-status-url]
[![Test status][test-status-shield]][test-status-url]
[![MIT License][license-shield]][license-url]

<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/AfoninaOlga/host-configurator.svg?style=for-the-badge&color=blue
[license-url]: LICENSE
[status-shield]: https://img.shields.io/github/actions/workflow/status/AfoninaOlga/host-configurator/.github/workflows/build.yml?branch=main&event=push&style=for-the-badge
[status-url]: https://github.com/AfoninaOlga/host-configurator/blob/main/.github/workflows/build.yml
[linter-status-shield]: https://img.shields.io/github/actions/workflow/status/AfoninaOlga/host-configurator/.github/workflows/lint.yml?branch=main&event=push&style=for-the-badge&label=Lint
[linter-status-url]: https://github.com/AfoninaOlga/host-configurator/blob/main/.github/workflows/lint.yml
[test-status-shield]: https://img.shields.io/github/actions/workflow/status/AfoninaOlga/host-configurator/.github/workflows/test.yml?branch=main&event=push&style=for-the-badge&label=Tests
[test-status-url]: https://github.com/AfoninaOlga/host-configurator/blob/main/.github/workflows/test.yml

# Host Configurator

gRPC-service changing hostname and updating DNS servers list.

It allows
- Get and set hostname
- Get DNS servers list
- Add DNS server
- Remove DNS server

## Dependencies

- Protocol Buffer Compiler &mdash; [protoc](https://grpc.io/docs/protoc-installation/)
- protocol compiler plugins for Go
```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Optional

- Make (for convenient building)

## Installation

Make sure `$PATH` variable includes `$HOME/go/bin` path.
If it does not, `export PATH="$PATH:$HOME/go/bin"`

### Service
```bash
make -C service
```

### Client

```bash
make -C client
```

## Run

### Run service

```bash
cd service
```

#### Explicit run

##### Run service


Replace `grpc_addr: "hostname-service"` with `grpc_addr: "127.0.0.1"` in `configs/gateway.yaml`

Service should be run as root to have permission to write `/etc/hostname` and `/etc/resolv.conf`

```bash
sudo ./service
```

##### Run gRPC-gateway

```bash
./gateway
```

#### Run with docker-compose

```bash
docker compose up -d --build
```

### Run client

```bash
cd client
```

#### Get hostname

```bash
./configurator get-hostname
```

#### Set hostname

```bash
./configurator set-hostname <hostname>
```

*hostname* validity is checked on the basis of matching the description from the [man](https://man7.org/linux/man-pages/man7/hostname.7.html)
>  Each element of the hostname must be from 1 to 63 characters long
and the entire hostname, including the dots, can be at most 253
characters long.  Valid characters for hostnames are ASCII(7)
letters from a to z, the digits from 0 to 9, and the hyphen (-).
A hostname may not start with a hyphen.
>
> -- <cite>Linux manual page</cite>

#### Get DNS servers list

```bash
./configurator dns-servers-list
```

#### Add DNS server

```bash
./configurator dns-servers-add <address>
```

### Remove DNS server

```bash
./configurator dns-servers-delete <address>
```

## Test

### End-to-end test

Before run make sure `grpc_addr` is set to `"hostname-service"` in `service/configs/gateway.yaml`

```bash
./e2e.sh
```
### Unit tests

TODO
