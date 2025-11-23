# Go Kafka Sender

A simple command-line utility to send messages from a file to Apache Kafka.

## Overview

Go Kafka Sender is a lightweight tool that reads messages from a text file and sends each line as a separate message to a specified Kafka topic. It's designed for batch processing of messages and supports flexible configuration through command-line arguments and configuration files.

## Features

- Send messages from a text file to Kafka
- Support for command-line arguments
- Configuration via YAML files
- Default configuration values

## Installation

### Prerequisites

- Go 1.25.3 or higher
- Apache Kafka cluster

### Build from Source

```bash
git clone https://github.com/GeorgeKuzora/go-kafka-sender.git
cd go-kafka-sender
go build -o go-kafka-sender ./cmd
```

## Usage

### Basic Usage

```bash
./go-kafka-sender <file-path>
```

### With Topic

```bash
./go-kafka-sender <topic> <file-path>
```

### With URL and Topic

```bash
./go-kafka-sender <url> <topic> <file-path>
```

### Arguments

- `url`: Kafka broker URL (default: http://127.0.0.1:9092)
- `topic`: Kafka topic name
- `file-path`: Path to the text file containing messages (one per line)

## Configuration

The application supports configuration through multiple sources in the following order of precedence:

1. Command-line arguments
2. User configuration file (`~/.config/gokafka/config.yaml`)
3. System configuration (not implemented yet)

### Configuration File Format

Create a configuration file at `~/.config/gokafka/config.yaml`:

```yaml
url: "http://localhost:9092"
topic: "default-topic"
```

## Input File Format

The input file should contain one message per line. Each line will be sent as a separate Kafka message.

Example:
```
Message 1
Message 2
Message 3
```

## Example

Send messages from a file to a Kafka topic:

```bash
./go-kafka-sender my-topic messages.txt
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
