### Why does this exist

The intention of this repo is just as a skeleton from various sources that I've been learning from. This is highly inspired by https://github.com/golang-standards/project-layout

### Why the name gofy?

Well, I couldn't think of anything else. gofy = go funk yourself

### What concepts does this cover

1. [Initialising application with cobra](old_stuff/cmd/main.go)
2. [Cobra sub command, ex: api](old_stuff/cmd/api/api.go)
3. [Initialise DB, ex: MongoDB](old_stuff/internal/db/db.go)
4. Using channels
5. Using go routines
6. [Implementing a worker pool](old_stuff/cmd/workers/workers.go)
7. Creating a kafka producer and consumer
8. gRPC in golang
9. Using protobuf
10. Dockerizing a go app
11. Logging
12. Managing configs
13. Concurrency
14. [Weather Application with Go Templates and CI/CD](web/)

---

## 🌦️ Featured Project: Weather Application

[![Docker Build](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml)
[![Docker Image Version](https://img.shields.io/docker/v/nikhildev/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Pulls](https://img.shields.io/docker/pulls/nikhildev/gofy-weather?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Image Size](https://img.shields.io/docker/image-size/nikhildev/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)

A beautiful weather application with Apple-inspired glassmorphism design, featuring:

- 🌡️ Real-time weather data from Open-Meteo API
- 📊 Interactive charts with Chart.js
- 🗺️ Global location search with country flags
- 📱 Responsive design for all devices
- 🐳 Production-ready Docker setup
- 🔄 GitHub Actions CI/CD pipeline
- ⚡ Hot reload development with Air

### Quick Start

```bash
# Run with Docker
docker run -d -p 5050:8080 nikhildev/gofy-weather:latest

# Or with docker-compose
cd web && docker-compose up -d

# Visit
open http://localhost:5050
```

**[📖 Full Documentation →](web/README.md)**
