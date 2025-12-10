# README Badges

Add these badges to your `README.md` once you have set up automated builds.

## Docker Hub

Ready-to-use badges for Docker Hub username: **nikhildev**

```markdown
[![Docker Image Version](https://img.shields.io/docker/v/nikhildev/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Image Size](https://img.shields.io/docker/image-size/nikhildev/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Pulls](https://img.shields.io/docker/pulls/nikhildev/gofy-weather?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
```

## GitHub Actions

Ready-to-use badge for GitHub repository (update if your repo name differs):

```markdown
[![Docker Build](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml)
```

## GitHub Container Registry

```markdown
[![GitHub Container Registry](https://img.shields.io/badge/ghcr.io-gofy--weather-blue?logo=github)](https://github.com/nikhildev/gofy/pkgs/container/gofy-weather)
```

## Example Usage in README

```markdown
# Gofy Weather Application

[![Docker Build](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml)
[![Docker Image Version](https://img.shields.io/docker/v/nikhildev/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Pulls](https://img.shields.io/docker/pulls/nikhildev/gofy-weather?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)
[![Docker Image Size](https://img.shields.io/docker/image-size/nikhildev/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)

A beautiful weather application built with Go...

## Quick Start

```bash
docker run -p 5050:8080 nikhildev/gofy-weather:latest
```

Visit http://localhost:5050
```

## All Available Badges

| Badge | Code |
|-------|------|
| Build Status | `[![Docker Build](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml/badge.svg)](https://github.com/nikhildev/gofy/actions/workflows/docker-build.yml)` |
| Docker Version | `[![Docker Image Version](https://img.shields.io/docker/v/nikhildev/gofy-weather?sort=semver&logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)` |
| Docker Pulls | `[![Docker Pulls](https://img.shields.io/docker/pulls/nikhildev/gofy-weather?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)` |
| Docker Size | `[![Docker Image Size](https://img.shields.io/docker/image-size/nikhildev/gofy-weather/latest?logo=docker)](https://hub.docker.com/r/nikhildev/gofy-weather)` |
| License | `[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)` |
| Go Version | `[![Go Version](https://img.shields.io/github/go-mod/go-version/nikhildev/gofy)](https://golang.org/)` |

---

**Note**: Update the repository name in badges if your GitHub repo isn't `nikhildev/gofy`. Badges will work after:
1. Setting up GitHub Actions (see [DOCKER_SETUP.md](DOCKER_SETUP.md))
2. First successful build and push to Docker Hub

