#!/bin/bash
# Quick setup script for development

set -e

echo "🚀 Starting Absensi Sekolah Backend Setup..."

# Check if .env exists
if [ ! -f .env ]; then
    echo "📝 Creating .env from .env.example..."
    cp .env.example .env
    echo "✅ .env created. Please update it with your configuration."
fi

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

echo "🐳 Starting Docker containers..."
docker-compose up -d

echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 5

echo "✅ Backend is running at http://localhost:8080"
echo ""
echo "Available endpoints:"
echo "  • POST   /api/v1/auth/login"
echo "  • POST   /api/v1/scan/tap"
echo "  • GET    /api/v1/reports/absensi"
echo ""
echo "To stop: docker-compose down"
echo "To view logs: docker-compose logs -f backend"
