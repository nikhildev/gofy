# Gofy Weather Application

A beautiful, modern weather application built with Go, featuring Apple-inspired glassmorphism design with real-time weather data and interactive charts. The application serves directly on the root path, providing instant access to weather information.

## 🌟 Features

### Weather Display
- **Current Conditions**: Real-time temperature, weather condition, and location with country flags
- **Today's Summary**: High/Low temperatures, sunrise/sunset times, and precipitation
- **48-Hour Forecast**: Detailed hourly forecast with temperature and precipitation charts
- **10-Day Forecast**: Temperature trends chart showing min/max temperatures
- **Daily Forecast List**: Clickable daily items with detailed modal view
- **Country Flags**: SVG flag icons for all locations using flag-icons library

### Interactive Features
- **Location Search**: Real-time city search with autocomplete and country flags
- **Clickable Day Details**: Click any day in the forecast to see 24-hour detailed view
- **Responsive Charts**: Interactive temperature and precipitation graphs using Chart.js
- **Hot Reload**: Automatic server restart on code changes using Air

### Design
- **Glassmorphism UI**: Apple-inspired frosted glass design
- **Responsive Layout**: Optimized for desktop, tablet, and mobile devices
- **Animated Backgrounds**: Dynamic gradient with floating blobs
- **Modern Typography**: Dancing Script handwriting font for titles

## 📋 Prerequisites

- Go 1.24.0 or higher
- Air (for hot reloading during development)

## 🚀 Setup Instructions

### 1. Install Dependencies

The project uses Go modules for dependency management. All required dependencies will be automatically downloaded when you build or run the project.

```bash
cd /Users/dev/code/ndc/gofy/web
go mod tidy
```

### 2. Install Air (Optional, for Development)

Air provides live reload functionality during development:

```bash
go install github.com/air-verse/air@latest
```

Make sure `$GOPATH/bin` is in your PATH, or use the full path to the Air binary.

## 🏃 Running the Application

### Option 1: Using Air (Recommended for Development)

Air automatically rebuilds and restarts the server when you make changes:

```bash
cd /Users/dev/code/ndc/gofy/web
air
# or with full path
/Users/dev/go/bin/air
```

### Option 2: Using Go Run

For a simple run without hot reload:

```bash
cd /Users/dev/code/ndc/gofy/web
go run main.go
```

### Option 3: Build and Run

Build the binary first, then run it:

```bash
cd /Users/dev/code/ndc/gofy/web
go build -o weather-app
./weather-app
```

The application will start on **http://localhost:8080**

## 🐳 Running with Docker

### Pre-built Images (Coming Soon)

Once set up, you can pull and run pre-built images:

```bash
# From Docker Hub
docker pull yourusername/gofy-weather:latest
docker run -d -p 8080:8080 yourusername/gofy-weather:latest

# From GitHub Container Registry
docker pull ghcr.io/yourusername/gofy-weather:latest
docker run -d -p 8080:8080 ghcr.io/yourusername/gofy-weather:latest
```

**Setup automated builds**: See [Docker Setup Guide](../../.github/DOCKER_SETUP.md) for configuring GitHub Actions to automatically build and push images.

### Quick Start with Docker (Local Build)

The easiest way to run the application is using Docker:

```bash
# Build the Docker image
cd /Users/dev/code/ndc/gofy/web
./build-docker.sh

# Run the container
docker run -d -p 8080:8080 --name gofy-weather gofy-weather:latest

# Access the application at http://localhost:8080
```

### Building the Docker Image

Use the provided build script:

```bash
# Basic build
./build-docker.sh

# Build with custom tag
./build-docker.sh --tag v1.0.0

# Build with custom name
./build-docker.sh --name my-weather-app

# Build and specify registry for pushing
./build-docker.sh --tag v1.0.0 --registry docker.io/username
```

### Build Script Options

| Option | Description | Example |
|--------|-------------|---------|
| `-t, --tag` | Set image tag | `--tag v1.0.0` |
| `-n, --name` | Set image name | `--name my-weather-app` |
| `-r, --registry` | Set registry URL | `--registry docker.io/myuser` |
| `-h, --help` | Show help message | `--help` |

### Manual Docker Build

If you prefer to build manually (build from parent directory):

```bash
cd /Users/dev/code/ndc/gofy
docker build -f web/Dockerfile -t gofy-weather:latest .
```

**Note**: The Docker build context must be the project root directory (not the `web` directory) because the Dockerfile needs access to `go.mod` and `go.sum` from the parent directory.

### Running the Container

```bash
# Run in detached mode
docker run -d -p 8080:8080 --name gofy-weather gofy-weather:latest

# Run with custom port (e.g., 3000)
docker run -d -p 3000:8080 --name gofy-weather gofy-weather:latest

# Run in foreground with logs
docker run -p 8080:8080 --name gofy-weather gofy-weather:latest

# Run with restart policy
docker run -d -p 8080:8080 --restart unless-stopped --name gofy-weather gofy-weather:latest
```

### Docker Container Management

```bash
# View logs
docker logs -f gofy-weather

# Check container status
docker ps | grep gofy-weather

# Stop the container
docker stop gofy-weather

# Start the container
docker start gofy-weather

# Restart the container
docker restart gofy-weather

# Remove the container
docker rm gofy-weather

# Remove the image
docker rmi gofy-weather:latest
```

### Docker Image Features

- **Multi-stage build**: Optimized for minimal image size
- **Non-root user**: Runs as unprivileged user for security
- **Health check**: Built-in health monitoring
- **Alpine-based**: Small footprint (~20-30 MB)
- **CA certificates**: Included for HTTPS API calls
- **Timezone data**: Included for accurate time handling

### Docker Compose (Recommended)

A `docker-compose.yml` file is included for easy deployment. Run from the `web` directory:

```bash
cd /Users/dev/code/ndc/gofy/web
docker-compose up -d        # Build and start in background
docker-compose logs -f      # View logs
docker-compose down         # Stop and remove containers
docker-compose restart      # Restart the service
```

## 🌐 Endpoints

### Web Pages

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Main weather application (default: Espoo, Finland) |
| `/?location={city}` | GET | Weather for a specific city (e.g., `/?location=London`) |

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/search-location?q={query}` | GET | Search for cities (returns up to 5 results with country codes) |

## 📁 Project Structure

```
gofy/
├── .dockerignore          # Docker build exclusions (parent directory)
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies
└── web/
    ├── main.go                 # Application entry point and routing
    ├── Dockerfile             # Docker image definition
    ├── docker-compose.yml     # Docker Compose configuration
    ├── build-docker.sh        # Docker build script
    ├── .air.toml              # Air configuration for hot reload
    ├── handlers/
    │   └── weather_handler.go # Weather page and API handlers
    ├── templates/
    │   └── weather.tmpl       # Weather application template
    └── tmp/                   # Air build artifacts (auto-generated)
```

**Note**: The `.dockerignore` file is located in the project root directory (`gofy/`) because the Docker build context is the parent directory.

## 🎨 Technology Stack

### Backend
- **Go**: Primary language
- **net/http**: HTTP server
- **html/template**: Server-side templating

### Frontend
- **HTML5/CSS3**: Modern semantic markup and styling
- **JavaScript (ES6+)**: Interactive features
- **Chart.js**: Interactive charts and graphs
- **flag-icons**: SVG country flag library

### External APIs
- **Open-Meteo API**: Free weather data (no API key required)
  - Weather Forecast: `https://api.open-meteo.com/v1/forecast`
  - Geocoding: `https://geocoding-api.open-meteo.com/v1/search`

## 🔧 Configuration

### Air Configuration (`.air.toml`)

The Air configuration is already set up to:
- Watch `.go`, `.tmpl`, and `.html` files
- Exclude `tmp/`, `vendor/`, and test files
- Build to `tmp/main` directory
- Auto-restart on file changes

### Port Configuration

The server runs on port `8080` by default. To change it, modify `main.go`:

```go
http.ListenAndServe(":8080", nil) // Change 8080 to your desired port
```

## 📱 Responsive Breakpoints

- **Desktop (≥1024px)**: 2x2 grid layout with side-by-side widgets
- **Tablet (768px-1023px)**: Stacked layout with adjusted padding
- **Mobile (480px-767px)**: Optimized for touch with larger buttons
- **Small Mobile (<480px)**: Compact layout with minimal padding

## 🎯 Features In Detail

### Weather Data
- Current temperature and conditions
- Today's high/low temperatures
- Sunrise and sunset times
- Precipitation levels
- 48-hour detailed forecast
- 10-day temperature trends
- Hourly breakdown for any selected day

### Interactive Charts
- **48-Hour Chart**: Line chart for temperature, bar chart for precipitation
- **10-Day Chart**: Min/max temperature trends
- **Daily Detail Chart**: 24-hour temperature and precipitation for selected day
- All charts are responsive and interactive with hover tooltips

### Search Functionality
- Real-time city search with debouncing (300ms)
- Autocomplete with up to 5 suggestions
- Country flags displayed for each result
- Click to navigate to selected city's weather

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Find process using port 8080
lsof -ti:8080

# Kill the process
kill -9 $(lsof -ti:8080)

# For Docker containers
docker stop gofy-weather
```

### Air Not Found
Make sure Air is installed and in your PATH:
```bash
# Check if Air is installed
which air

# If not in PATH, use full path
/Users/dev/go/bin/air
```

### Templates Not Found
Make sure you're running the application from the `web` directory:
```bash
cd /Users/dev/code/ndc/gofy/web
go run main.go
```

### Docker Build Fails
```bash
# Clean up Docker build cache
docker builder prune

# Rebuild without cache
docker build --no-cache -t gofy-weather:latest .

# Check Docker daemon is running
docker ps
```

### Container Won't Start
```bash
# Check container logs
docker logs gofy-weather

# Check if port is already in use
docker ps | grep 8080

# Try a different port
docker run -d -p 3000:8080 --name gofy-weather gofy-weather:latest

# Test if the app is accessible
curl http://localhost:8080/?location=Espoo
```

### Health Check Failing
```bash
# Check if the app is responding
docker exec gofy-weather wget -O- http://localhost:8080/?location=Espoo

# Check health status
docker inspect --format='{{.State.Health.Status}}' gofy-weather

# View detailed health check logs
docker inspect --format='{{json .State.Health}}' gofy-weather | python3 -m json.tool
```

## 📝 Development Tips

1. **Hot Reload**: Use Air during development for instant feedback
2. **Template Changes**: Templates reload automatically with Air
3. **Handler Changes**: Go files trigger rebuild and restart
4. **CSS Changes**: Template file changes reload the server
5. **Browser DevTools**: Use browser console to debug JavaScript

## 🌍 Supported Locations

The application supports weather data for **any location worldwide** using the Open-Meteo API's geocoding service. Simply search for any city, town, or location name.

**Examples:**
- Default: `http://localhost:8080/` (shows Espoo, Finland)
- London: `http://localhost:8080/?location=London`
- New York: `http://localhost:8080/?location=New York`
- Tokyo: `http://localhost:8080/?location=Tokyo`

## 📦 Dependencies

All dependencies are managed via `go.mod`:

```go
require (
    // No external Go dependencies required!
    // Uses only Go standard library
)
```

Frontend libraries loaded via CDN:
- Chart.js v4.x
- flag-icons v7.2.3
- Google Fonts (Inter, Dancing Script)

## 🤖 CI/CD with GitHub Actions

The project includes automated Docker builds using GitHub Actions:

### Workflows

1. **Automatic Builds** (`.github/workflows/docker-build.yml`)
   - Triggers on push to main/master
   - Builds for AMD64 and ARM64 platforms
   - Pushes to Docker Hub and GitHub Container Registry
   - Creates version tags from GitHub releases

2. **Manual Builds** (`.github/workflows/docker-build-manual.yml`)
   - Manually trigger from Actions tab
   - Choose custom tags
   - Select registries to push to

### Setup Instructions

See the complete guide: [Docker Setup Guide](../../.github/DOCKER_SETUP.md)

**Quick Setup:**
1. Create Docker Hub account and access token
2. Add `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` to GitHub secrets
3. Push to main branch or create a release
4. Images automatically built and published!

## 🤝 Contributing

To contribute to this project:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly on multiple devices
5. Ensure Air hot reload works correctly
6. Check responsive design on different screen sizes
7. Create a pull request (Docker images will be built automatically for testing)

## 📄 License

Part of the Gofy project - a Go learning repository.

## 🔗 Related Projects

This weather app is part of the larger Gofy monorepo which includes various Go examples and projects.

---

**Built with ❤️ using Go and modern web technologies**
