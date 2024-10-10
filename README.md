# Redis Clone

## Overview

This project is a lightweight in-memory key-value store developed in Go, designed to replicate essential features of Redis. It supports various data structures and provides a scalable architecture for concurrent client interactions.

## Command Implementation Checklist

#### Implemented Commands

- [x] **SET** key value
- [x] **GET** key
- [x] **HSET** key field value
- [x] **HGET** key field
- [x] **HGETALL** key

#### Unimplemented Commands

- [ ] **EXPIRE** key seconds
- [ ] **DEL** key
- [ ] **EXISTS** key
- [ ] **LPUSH** key value
- [ ] **RPUSH** key value

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

- Go 1.23 or later

### Local Build Instructions

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
./dist/redis-clone
```

### Docker Setup

If you prefer to run the project in a Docker container, follow these instructions:

1. Clone the repository (if you haven't already):

```bash
git clone https://github.com/ghulamazad/redis-clone.git
cd redis-clone
```

2. Build the Docker image:

```bash
docker build -t redis-clone .
```

3. Run the Docker container:

```bash
docker run -p 6379:6379 redis-clone
```

This command maps port 6379 of your host machine to port 6379 of the container. Adjust the port number if your application uses a different one.

### Stopping the Container

To stop the running container, find its container ID with:

```bash
docker ps
```

Then use the following command to stop it:

```bash
docker stop <container_id>
```

Replace <container_id> with the actual ID of the running container.

### Cleaning Up

If you want to remove the stopped container and the image, use:

```bash
docker rm <container_id>
docker rmi redis-clone
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
