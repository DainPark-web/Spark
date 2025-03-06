# Spark

I built Spark as a fun way to learn Go while making something useful. It started as a basic web server experiment but grew into a full dating app. I focused on using Go's strengths - like fast performance and easy concurrency - to create a reliable platform where people can meet and connect with others.

## Project Overview

Spark offers the following features:
- RESTful API server built with Gin framework
- User authentication and profile management
- CORS configuration for secure cross-origin requests
- Environment-based configuration for development and production
- Trusted proxy support for production deployments
- Database integration for user data persistence
- Structured error handling and status responses


## Installation and Running
```bash
cd/code/server

# Downloads and installs all required dependencies
go mod tidy
```
### Run server using air
```bash
# run server
# if you are using air
air init
air
```
```toml
cmd = "go build -o ./tmp/main ./cmd"
```

### Run server using go
```bash
go run cmd/main.go
```

## Contributing
- coming soon...
## License
- coming soon...

