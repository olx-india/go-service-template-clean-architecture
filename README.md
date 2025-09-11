# Go Service Template—Clean Architecture

A production-ready Go service template following Clean Architecture principles, designed for microservices with comprehensive infrastructure support.

### Maintainers:

| Manager                | Developer              |
|------------------------|------------------------|
| `bhajan.poonia@olx.in` | `parvez.hassan@olx.in` |

### Protected Branches

| Branches  |
|-----------|
| `main`    |
| `develop` |

### Branch Naming Convention

| Value                        | Description                                             |
|------------------------------|---------------------------------------------------------|
| `develop`                    | Stable branch to create feature branches from           |
| `feature/OLXIN-0000-feature` | Feature branch created after an assigned JIRA ticket Id |
| `release/0.0.0`              | Release candidate for staging sanity                    |
| `main`                       | Deploy to Production                                    |

### Branch Merge Flow

Normal: `feature/OLXIN-0000-feature` -> `develop` -> `release/0.0.0` -> `main`

Hotfix: `hotfix/OLXIN-000-hot-fix` -> `main`, then take a pull of the changes into `release/0.0.0` & `develop` branches

### CI/CD Pipeline Jobs

| `Stage`           | `Work`                                                                |
|-------------------|-----------------------------------------------------------------------|
| `build`           | Uses build tools to create application jar or compile code            |
| `test`            | Integration of Unit Tests and Integration Tests                       |
| `quality`         | Stage for running Sonarqube tests and upload reports and score        |
| `package`         | Create a docker image that can be deployed to the staging environment |
| `terraform-plan`  | Plan stage to preview infrastructure changes                          |
| `terraform-apply` | Apply infrastructure changes on the cloud                             |
| `helm-lint`       | Lint check resources. List changes in resources to be deployed        |
| `helm-deploy`     | Deploy stage to rollout application changes                           |
| `helm-rollback`   | Rollback stage in case changes need to be reverted to last stable     |
| `blank-rollout`   | Restart the last deployed state of the application                    |

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
- **New Relic Monitoring**: Application performance monitoring
- **Docker Support**: Containerized deployment with multi-stage builds
- **Terraform Infrastructure**: AWS infrastructure as code
- **Health Checks**: Kubernetes-ready health endpoints
- **Comprehensive Testing**: Unit and integration test support
- **Code Quality**: Linting, formatting, and security scanning

## 📋 Prerequisites

- Go 1.24.0 or higher
- Docker and Docker Compose
- Make
- Terraform (for infrastructure deployment)

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
| `NEW_RELIC_LICENSE_KEY` | New Relic license key | `default-license-key` |
| `READ_TIMEOUT` | HTTP read timeout | `60s` |
| `WRITE_TIMEOUT` | HTTP write timeout | `60s` |

## 🏗️ Infrastructure

### Terraform Deployment

The project includes Terraform configurations for AWS infrastructure:

- **RDS Aurora PostgreSQL**: Managed database cluster
- **ElastiCache Redis**: Managed Redis cluster
- **Route53**: DNS management
- **Security Groups**: Network security
- **IAM Roles**: Access management

#### Deploy Infrastructure

1. **Navigate to terraform directory**
   ```bash
   cd terraform
   ```

2. **Initialize Terraform**
   ```bash
   terraform init
   ```

3. **Plan deployment**
   ```bash
   terraform plan -var-file="vars/qa/terraform.tfvars"
   ```

4. **Apply configuration**
   ```bash
   terraform apply -var-file="vars/qa/terraform.tfvars"
   ```

### Environment-Specific Configuration

The project supports multiple environments:
- `stg/` - Staging environment
- `qa/` - QA environment  
- `prd/` - Production environment

Each environment has its own variable files in the `vars/` directory.

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
├── terraform/             # Infrastructure as code
├── vars/                  # Environment configurations
└── docker-compose.yml     # Local development setup
```

### Adding New Features

1. **Define domain entities** in `internal/domain/`
2. **Implement use cases** in `internal/usecase/`
3. **Create API handlers** in `internal/api/`
4. **Add routes** in `server/router/`
5. **Write tests** for all layers

### Database Migrations

```bash
# Create new migration
make migrate-create migration_name

# Apply migrations
make migrate-up
```

## 🚀 Deployment

### Docker
```bash
# Build image
docker build -t go-service-template .

# Run container
docker run -p 8080:8080 go-service-template
```

### Kubernetes
The service is designed to be deployed on Kubernetes with:
- Health check endpoints for liveness/readiness probes
- Graceful shutdown handling
- Configurable resource limits
- Horizontal Pod Autoscaling support

## 📊 Monitoring

### New Relic Integration
The service includes New Relic APM integration for:
- Application performance monitoring
- Error tracking
- Custom metrics
- Distributed tracing

### Logging
Structured logging with:
- JSON format for production
- Context propagation
- Correlation IDs
- Log levels (DEBUG, INFO, WARN, ERROR)

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
