# Docker Hub Setup Guide

This guide will help you set up automated Docker image builds and pushes to Docker Hub and GitHub Container Registry using GitHub Actions.

## 📋 Prerequisites

1. **Docker Hub Account**
   - Sign up at https://hub.docker.com
   - Verify your email

2. **GitHub Repository**
   - Your code must be in a GitHub repository
   - You need admin access to the repository

## 🔑 Step 1: Create Docker Hub Access Token

1. Go to https://hub.docker.com
2. Click on your username (top right) → **Account Settings**
3. Click on **Security** → **New Access Token**
4. Token description: `github-actions-gofy`
5. Access permissions: **Read, Write, Delete**
6. Click **Generate**
7. **Copy the token** (you won't see it again!)

## 🔐 Step 2: Add Secrets to GitHub Repository

1. Go to your GitHub repository
2. Click **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**

Add these two secrets:

### Secret 1: DOCKERHUB_USERNAME
- Name: `DOCKERHUB_USERNAME`
- Value: `nikhildev`

### Secret 2: DOCKERHUB_TOKEN
- Name: `DOCKERHUB_TOKEN`
- Value: The access token you copied from Docker Hub

## ✅ Step 3: Verify Setup

The workflow will automatically run when you:
- Push to `main` or `master` branch
- Create a pull request
- Create a release
- Manually trigger the workflow

## 🚀 Usage

### Automatic Builds

The workflow automatically builds and pushes images when you push code:

```bash
git add .
git commit -m "Update weather app"
git push origin main
```

This will:
- Build the Docker image for AMD64 platform
- Push to Docker Hub: `nikhildev/gofy-weather:latest`
- Push to GHCR: `ghcr.io/nikhildev/gofy-weather:latest`

### Manual Builds

1. Go to your GitHub repository
2. Click **Actions** tab
3. Click **Manual Docker Build** workflow
4. Click **Run workflow**
5. Choose:
   - Tag (e.g., `v1.0.0`, `latest`)
   - Push to Docker Hub (yes/no)
   - Push to GHCR (yes/no)
6. Click **Run workflow**

### Release Builds

Create a GitHub release:

```bash
# Create and push a tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0

# Or create release via GitHub UI
# Go to: Releases → Create a new release
```

This will create images with tags:
- `nikhildev/gofy-weather:v1.0.0`
- `nikhildev/gofy-weather:1.0`
- `nikhildev/gofy-weather:1`
- `nikhildev/gofy-weather:latest`

## 📦 Pulling Your Image

Once published, anyone can pull your image:

```bash
# From Docker Hub
docker pull nikhildev/gofy-weather:latest

# From GitHub Container Registry
docker pull ghcr.io/nikhildev/gofy-weather:latest

# Run the container
docker run -d -p 5050:8080 nikhildev/gofy-weather:latest
# Or with docker-compose: docker-compose up -d
```

## 🔍 Monitoring Builds

### View Build Status
1. Go to **Actions** tab in your GitHub repository
2. Click on the workflow run
3. See build logs and status

### Check Published Images

**Docker Hub:**
- Visit: `https://hub.docker.com/r/nikhildev/gofy-weather`

**GitHub Container Registry:**
- Visit: `https://github.com/nikhildev/gofy/pkgs/container/gofy-weather`

## 🎯 Workflow Features

✅ **Multi-platform builds**: AMD64 and ARM64 (works on Intel, M1/M2 Macs, ARM servers)
✅ **Build caching**: Faster subsequent builds using GitHub Actions cache
✅ **Automatic tagging**: Smart tagging based on branches, tags, and commits
✅ **Security**: Uses access tokens instead of passwords
✅ **Dual registry**: Pushes to both Docker Hub and GitHub Container Registry
✅ **Pull request testing**: Builds (but doesn't push) on PRs

## 🛠️ Troubleshooting

### Build Fails: "secrets.DOCKERHUB_USERNAME not found"
- Make sure you added the secrets in GitHub repository settings
- Secret names must match exactly: `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN`

### Build Fails: "denied: requested access to the resource is denied"
- Check your Docker Hub access token has write permissions
- Verify your Docker Hub username is correct
- Regenerate the access token if needed

### Image Not Appearing on Docker Hub
- Check if the workflow ran successfully in Actions tab
- Verify you're logged into Docker Hub with the correct account
- Image might be private by default - change visibility in Docker Hub settings

## 📝 Advanced Configuration

### Change Image Name

Edit `.github/workflows/docker-build.yml`:
```yaml
env:
  IMAGE_NAME: your-custom-name  # Change this
```

### Add More Platforms

Edit the `platforms` line:
```yaml
platforms: linux/amd64,linux/arm64,linux/arm/v7
```

### Disable GitHub Container Registry

Remove the GHCR login and build steps from the workflow file.

## 🔗 Useful Links

- Docker Hub: https://hub.docker.com
- GitHub Actions Documentation: https://docs.github.com/en/actions
- Docker Build Push Action: https://github.com/docker/build-push-action
- GitHub Container Registry: https://docs.github.com/en/packages

## 💡 Tips

1. **Keep tokens secure**: Never commit tokens to your repository
2. **Use releases**: Create GitHub releases for version tags
3. **Monitor builds**: Check Actions tab regularly for failed builds
4. **Update README**: Add Docker Hub badges to your README

---

**Need help?** Check the Actions tab for detailed build logs and error messages.

