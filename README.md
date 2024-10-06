# Redis Clone

## Overview

This project is a lightweight in-memory key-value store developed in Go, designed to replicate essential features of Redis. It supports various data structures and provides a scalable architecture for concurrent client interactions.

## Features

- **Data Structures**: Implemented support for:

  - Strings
  - Hashes
  - Lists
  - Sets
  - Sorted Sets

- **Concurrent Client Support**: Engineered a client-server model that allows multiple clients to connect and interact seamlessly.
- **Redis Command Compatibility**: Supports standard Redis commands, enabling users to perform data manipulation with existing Redis client software.

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/ghulamazad/redis-clone.git
cd redis-clone
```

2. Build the project:

```bash
make build
```

3. Run the server:

```bash
./redis-clone
```

### Usage

Connect to the server using a Redis client and start executing commands. Here are a few examples:

- Set a key:

```bash
SET key "value"
```

- Get a key:

```bash
GET key
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any feature requests or bug reports.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

Inspired by Redis, this project aims to deepen understanding of key-value store architecture and concurrency in Go.

Feel free to explore and contribute!
