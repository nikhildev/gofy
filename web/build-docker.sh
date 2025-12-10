#!/bin/bash

# Gofy Weather App - Docker Build Script
# This script builds an OCI-compliant Docker image for the weather application

set -e  # Exit on error

# Configuration
IMAGE_NAME="gofy-weather"
IMAGE_TAG="latest"
REGISTRY=""  # Set this if you want to push to a registry (e.g., "docker.io/username")

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Check if Docker is installed
check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker first."
        exit 1
    fi
    print_success "Docker is installed"
}

# Build the Docker image
build_image() {
    local full_image_name="${IMAGE_NAME}:${IMAGE_TAG}"
    
    if [ -n "$REGISTRY" ]; then
        full_image_name="${REGISTRY}/${full_image_name}"
    fi
    
    print_info "Building Docker image: ${full_image_name}"
    print_info "This may take a few minutes..."
    
    # Get the script directory (web directory)
    local script_dir="$(cd "$(dirname "$0")" && pwd)"
    
    # Build from parent directory with web as context
    cd "${script_dir}/.."
    
    print_info "Building from: $(pwd)"
    docker build -f web/Dockerfile -t "${full_image_name}" .
    
    if [ $? -eq 0 ]; then
        print_success "Docker image built successfully: ${full_image_name}"
        return 0
    else
        print_error "Failed to build Docker image"
        return 1
    fi
}

# Get image size
get_image_size() {
    local full_image_name="${IMAGE_NAME}:${IMAGE_TAG}"
    
    if [ -n "$REGISTRY" ]; then
        full_image_name="${REGISTRY}/${full_image_name}"
    fi
    
    local size=$(docker images "${full_image_name}" --format "{{.Size}}")
    print_info "Image size: ${size}"
}

# Display run instructions
show_run_instructions() {
    local full_image_name="${IMAGE_NAME}:${IMAGE_TAG}"
    
    if [ -n "$REGISTRY" ]; then
        full_image_name="${REGISTRY}/${full_image_name}"
    fi
    
    echo ""
    print_success "Build complete!"
    echo ""
    print_info "To run the container:"
    echo "  docker run -d -p 8080:8080 --name gofy-weather ${full_image_name}"
    echo ""
    print_info "To run with custom port (e.g., 3000):"
    echo "  docker run -d -p 3000:8080 --name gofy-weather ${full_image_name}"
    echo ""
    print_info "To run with docker-compose:"
    echo "  docker-compose -f docker-compose.local.yml up -d"
    echo ""
    print_info "To view logs:"
    echo "  docker logs -f gofy-weather"
    echo ""
    print_info "To stop the container:"
    echo "  docker stop gofy-weather"
    echo ""
    print_info "To remove the container:"
    echo "  docker rm gofy-weather"
    echo ""
    print_info "Access the weather application at: http://localhost:8080"
    print_info "Try different cities: http://localhost:8080/?location=London"
    echo ""
    print_info "To push to Docker Hub (nikhildev):"
    echo "  docker tag ${full_image_name} nikhildev/gofy-weather:latest"
    echo "  docker push nikhildev/gofy-weather:latest"
}

# Push to registry (optional)
push_image() {
    if [ -z "$REGISTRY" ]; then
        print_warning "No registry configured. Skipping push."
        return 0
    fi
    
    local full_image_name="${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
    
    read -p "Do you want to push the image to ${REGISTRY}? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_info "Pushing image to registry..."
        docker push "${full_image_name}"
        
        if [ $? -eq 0 ]; then
            print_success "Image pushed successfully"
        else
            print_error "Failed to push image"
            return 1
        fi
    fi
}

# Main execution
main() {
    echo "======================================"
    echo "  Gofy Weather App - Docker Build"
    echo "======================================"
    echo ""
    
    # Check prerequisites
    check_docker
    
    # Build the image
    build_image
    
    # Get image information
    get_image_size
    
    # Show run instructions
    show_run_instructions
    
    # Optional: Push to registry
    push_image
    
    echo ""
    print_success "All done! 🎉"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --tag|-t)
            IMAGE_TAG="$2"
            shift 2
            ;;
        --name|-n)
            IMAGE_NAME="$2"
            shift 2
            ;;
        --registry|-r)
            REGISTRY="$2"
            shift 2
            ;;
        --help|-h)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Options:"
            echo "  -t, --tag TAG         Set image tag (default: latest)"
            echo "  -n, --name NAME       Set image name (default: gofy-weather)"
            echo "  -r, --registry REG    Set registry URL (e.g., docker.io/username)"
            echo "  -h, --help           Show this help message"
            echo ""
            echo "Example:"
            echo "  $0 --tag v1.0.0 --registry docker.io/myuser"
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

# Run main function
main

