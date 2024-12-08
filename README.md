## Distributed Locking Service

This repository provides a distributed locking service implemented over HTTP and gRPC interfaces, enabling distributed systems to synchronize access to shared resources across multiple nodes or instances. The service offers mechanisms to acquire and release locks, ensuring mutual exclusion in a distributed environment, in a serialised fashion.

## Table of Contents

0. [Overview](#overview)
1. [Features](#features)
2. [Architecture](#architecture)
3. [Getting Started](#getting-started)
4. [Contributing](#contributing)
5. [License](#license)

## Overview<a name="overview"></a>

This distributed locking service allows clients to request locks for resources over an HTTP and gRPC interface. It handles locking in a fault-tolerant and scalable way, ensuring that only one client can hold a lock on a given resource at any time. All the incoming requests for the same key are processed serially to ensure the order of acquisition. Supports automatic lock expiry after a deadline in case of client crash.

The service is ideal for use cases in microservices architectures, cloud-based applications, or systems that require synchronization and exclusive access to resources.

## Features<a name="features"></a>

- **Distributed Locking**: Supports acquiring and releasing locks across multiple nodes in a distributed system.
- **HTTP Interface**: Simple HTTP interface to quickly get started.
- **gRPC Interface**: Fast, efficient communication via gRPC for high-performance locking.
- **Timeouts and Retry Logic**: Configurable lock timeouts and retries for handling transient issues.
- **Automatic Expiry**: Locks automatically expire after a configurable time to prevent deadlocks in case of client crashes.

## Architecture<a name="architecture"></a>

The service provides HTTP and gRPC for communication. The architecture is designed to be fault-tolerant and provides automatic lock expiration to avoid deadlocks.

### Key Components<a name="key-components"></a>

1. HTTP and gRPC Server: Provides the API for acquiring and releasing locks.
2. Lock Manager: Handles the internal logic of managing lock state (acquire, release, expire).

## Getting Started<a name="getting-started"></a>

### Prerequisites<a name="prerequisites"></a>

Before you begin, ensure you have the following installed:

1. Go (v1.23.2 or higher)
2. Protocol Buffers (protoc) compiler
3. Protobuf plugins like protoc-gen-go and protoc-gen-go-grpc

### Installation<a name="installation"></a>

1. Clone the repository:

```
git clone https://github.com/sgkandale/ShareLock.git
cd ShareLock
```

2. Install Go dependencies:

```
go mod tidy
```

3. Generate the gRPC code from the .proto files:

```
./generatePB.sh
```

4. Cross check config.yaml file one for all the values

### Running the Service<a name="running-the-service"></a>

1. Build the binary

```
go build ./cmd/sharelock
```

2. Start the server

```
./sharelock
```

### Using the Service<a name="using-the-service"></a>

Once the server is running, you can use the client to interact with the locking service. For examples on how to use the service, refer the examples directory inside `cmd`

### API Reference<a name="api-reference"></a>

Payload structures are in sharelock.proto file. Refer the same to build your own gRPC client or to use over HTTP.

## Contributing<a name="contributing"></a>

Contributions to improve this project are always welcomed. If you'd like to help, please fork the repository and create a pull request. Here are some ways you can contribute:

- Bug fixes and optimizations
- Enhancing the documentation
- Implementing new features or lock types

Please follow the standard GitHub workflow:

1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Make your changes.
4. Submit a pull request with a clear description of your changes.

## License<a name="license"></a>

This project is licensed under the MIT License - see the LICENSE file for details.
