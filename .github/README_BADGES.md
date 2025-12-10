# README Badges

Add these badges to your `README.md` once you have set up automated builds.

## Docker Hub

Replace `yourusername` with your Docker Hub username:

```markdown
[![Docker Image Version](https://img.shields.io/docker/v/yourusername/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)
[![Docker Image Size](https://img.shields.io/docker/image-size/yourusername/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)
[![Docker Pulls](https://img.shields.io/docker/pulls/yourusername/gofy-weather?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)
```

## GitHub Actions

Replace `yourusername/gofy` with your GitHub repository:

```markdown
[![Docker Build](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml)
```

## GitHub Container Registry

Replace `yourusername` with your GitHub username:

```markdown
[![GitHub Container Registry](https://img.shields.io/badge/ghcr.io-gofy--weather-blue?logo=github)](https://github.com/yourusername/gofy/pkgs/container/gofy-weather)
```

## Example Usage in README

```markdown
# Gofy Weather Application

[![Docker Build](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml)
[![Docker Image Version](https://img.shields.io/docker/v/yourusername/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)
[![Docker Pulls](https://img.shields.io/docker/pulls/yourusername/gofy-weather?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)
[![Docker Image Size](https://img.shields.io/docker/image-size/yourusername/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)

A beautiful weather application built with Go...

## Quick Start

```bash
docker run -p 8080:8080 yourusername/gofy-weather:latest
```

Visit http://localhost:8080
```

## All Available Badges

| Badge | Code |
|-------|------|
| Build Status | `[![Docker Build](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/yourusername/gofy/actions/workflows/docker-build.yml)` |
| Docker Version | `[![Docker Image Version](https://img.shields.io/docker/v/yourusername/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)` |
| Docker Pulls | `[![Docker Pulls](https://img.shields.io/docker/pulls/yourusername/gofy-weather?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)` |
| Docker Size | `[![Docker Image Size](https://img.shields.io/docker/image-size/yourusername/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/yourusername/gofy-weather)` |
| License | `[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)` |
| Go Version | `[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/gofy)](https://golang.org/)` |

---

**Note**: Badges will only work after:
1. Setting up GitHub Actions (see [DOCKER_SETUP.md](DOCKER_SETUP.md))
2. First successful build and push to Docker Hub
3. Replacing `yourusername` with your actual username

