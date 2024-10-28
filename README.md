# Hello World Container

A simple, informative container that displays server information, environment variables, and client details. Perfect for testing container deployments and environment configurations.

![Docker Pulls](https://img.shields.io/docker/pulls/phillarmonic/hello-world)
![Docker Image Size](https://img.shields.io/docker/image-size/phillarmonic/hello-world)

## Features

- 📡 Display container name and server IP
- 🌐 Show client IP address
- 🔍 List all environment variables (with automatic masking of sensitive data)
- 📋 Copy environment variables as JSON
- 🔒 Security-conscious (masks sensitive environment variables)

## Quick Start

```bash
docker run -d -p 8080:8080 phillarmonic/hello-world
```

Then open your browser at `http://localhost:8080`

## Usage with Docker Compose

Create a `docker-compose.yml`:

```yaml
services:
  webserver:
    image: phillarmonic/hello-world
    ports:
      - "80:80"
    environment:
      - SERVER_IP=192.168.1.100  # Optional: Set your server IP
      - CUSTOM_VAR=test_value    # Add any custom environment variables
```

Then run:

```bash
docker compose up -d
```

## Environment Variables

| Variable    | Description                  | Default     |
|-------------|------------------------------|-------------|
| SERVER_IP   | IP address of the server     | 0.0.0.0    |
| Any others  | Will be displayed in UI      | -          |

Note: Environment variables containing "key", "token", "secret", or "password" will be automatically masked in the UI.

## Building from Source

```bash
# Clone the repository
git clone https://github.com/phillarmonic/hello-world

# Build the image
docker build -t phillarmonic/hello-world .

# Run the container
docker run -d -p 80:80 phillarmonic/hello-world
```

## Security Features

- Automatic masking of sensitive environment variables
- No external dependencies
- Minimal base image

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details

## About

Made with ❤️ by [Phillarmonic Software](https://github.com/phillarmonic)