# Go Service Template—Clean Architecture

A production-ready Go service template following Clean Architecture principles, designed for microservices with comprehensive infrastructure support.

## 🏗️ Architecture

This project implements Clean Architecture with the following layers:

- **Domain Layer** (`internal/domain/`): Core business entities and rules
- **Use Case Layer** (`internal/usecase/`): Application business logic
- **Infrastructure Layer** (`internal/infrastructure/`): External concerns (config, logging, repositories)
- **API Layer** (`internal/api/`): HTTP handlers and DTOs
- **Server Layer** (`server/`): Application setup and routing

## 🚀 Features

- **Clean Architecture**: Separation of concerns with dependency inversion
- **Gin Web Framework**: High-performance HTTP web framework
- **Structured Logging**: Zap-based logging with context propagation
- **Redis Integration**: Caching and session management
- **Docker Support**: Containerized deployment with multi-stage builds
- **Health Checks**: Kubernetes-ready health endpoints
- **Comprehensive Testing**: Unit and integration test support
- **Code Quality**: Linting, formatting, and security scanning

## 🧹 About Clean Architecture

The Clean Code Blog by Robert C. Martin (Uncle Bob) - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

- ![clean-architecture.png](clean-architecture.png)

## 📋 Prerequisites

- Go 1.24.0 or higher
- Docker and Docker Compose
- Make

### Installing Go with GVM (Go Version Manager)

GVM allows you to easily install and manage multiple Go versions. Here's how to set it up:

#### 1. Install GVM

**On macOS:**
```bash
# Install GVM
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# Add GVM to your shell profile
echo '[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"' >> ~/.zshrc
source ~/.zshrc
```

#### 2. Install Go 1.24.0

```bash
# Install Go 1.24.0
gvm install go1.24.0

# Use Go 1.24.0 as default
gvm use go1.24.0 --default

# Verify installation
go version
```

#### 3. GVM Commands Reference

```bash
# List available Go versions
gvm listall

# Install a specific Go version
gvm install go1.23.0

# Switch to a different Go version
gvm use go1.23.0

# Set a Go version as default
gvm use go1.24.0 --default

# List installed Go versions
gvm list

# Uninstall a Go version
gvm uninstall go1.23.0
```

#### Alternative: Direct Go Installation

If you prefer not to use GVM, you can install Go directly:

1. Download Go from [https://golang.org/dl/](https://golang.org/dl/)
2. Follow the installation instructions for your operating system
3. Verify installation with `go version`

## 🛠️ Quick Start

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-service-template-clean-architecture
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Setup Env Variables**
   ```
   Copy .env.example -> .env 
   Update .env variables to run service on local
   ```

4. **Run the application**
   ```bash
   make run
   ```

The service will be available at `http://localhost:8080`

### Run With Docker 

1. **Build and run with Docker Compose**
   ```bash
   make compose-up
   ```

2. **Stop services**
   ```bash
   make compose-down
   ```

## 📚 API Documentation

### Health Check
```http
GET /health
```

Response:
```json
{
  "status": "ok",
  "service": "go-service-template"
}
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `HOST` | Server host | `0.0.0.0` |
| `PORT` | Server port | `8080` |
| `REDIS_HOST` | Redis host | `localhost` |
| `ENV` | Environment | `local` |
| `APP_NAME` | Application name | `go-service-template` |
| `READ_TIMEOUT` | HTTP read timeout | `60s` |
| `WRITE_TIMEOUT` | HTTP write timeout | `60s` |

## 🧪 Testing

### Run Tests
```bash
# Unit tests
make test

# Integration tests
make integration-test

# All tests with coverage
make test && make integration-test
```

### Code Quality
```bash
# Lint code
make lint

# Format code
make format

# Run all pre-commit checks
make pre-commit
```

### Mockery Integration

This service includes Testify and Mockery for mocking interfaces.

#### Installation and Configuration

```bash
  go install github.com/vektra/mockery/v3@v3.5.5
  # check version
  mockery version
```
or through homebrew
```bash
  brew install mockery
```

#### Usage

To generate mocks for all the interfaces.

```bash
  mockery
```

Note — please refer to the [Mockery Documentation](https://vektra.github.io/mockery/latest/installation/) for more information and specific configurations.

## 📦 Development

### Project Structure
```
├── cmd/                    # Application entry point
├── internal/               # Private application code
│   ├── api/               # HTTP handlers and DTOs
│   ├── domain/            # Business entities
│   ├── infrastructure/    # External dependencies
│   └── usecase/           # Business logic
├── server/                # Application setup
└── docker-compose.yml     # Local development setup
```

### Adding New Features

1. **Define domain entities** in `internal/domain/`
2. **Implement use cases** in `internal/usecase/`
3. **Create API handlers** in `internal/api/`
4. **Add routes** in `server/router/`
5. **Write tests** for all layers

## 🚀 Deployment

### Docker
```bash
# Build image
docker build -t go-service-template .

# Run container
docker run -p 8080:8080 go-service-template
```

### Logging
Structured logging with:
- JSON format for production
- Context propagation
- Correlation IDs
- Log levels (DEBUG, INFO, WARN, ERROR)

## OpenTelemetry Integration

This service includes OpenTelemetry for distributed tracing and observability.

### Configuration

Add to your `.env` file:

### Local Development

Run Jaeger for trace visualization:

```bash
docker run -d --name jaeger -p 16686:16686 -p 14250:14250 jaegertracing/all-in-one:latest
```

Visit `http://localhost:16686` to view traces.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License—see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the code examples

---

**Built with ❤️ using Go and Clean Architecture principles**
