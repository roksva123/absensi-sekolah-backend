@echo off
REM Quick setup script for development on Windows

setlocal enabledelayedexpansion

echo.
echo 🚀 Starting Absensi Sekolah Backend Setup...
echo.

REM Check if .env exists
if not exist .env (
    echo 📝 Creating .env from .env.example...
    copy .env.example .env
    echo ✅ .env created. Please update it with your configuration.
)

REM Check if Docker is installed
docker --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker is not installed. Please install Docker Desktop first.
    exit /b 1
)

echo 🐳 Starting Docker containers...
docker-compose up -d

echo ⏳ Waiting for PostgreSQL to be ready...
timeout /t 5 /nobreak

echo.
echo ✅ Backend is running at http://localhost:8080
echo.
echo Available endpoints:
echo   • POST   /api/v1/auth/login
echo   • POST   /api/v1/scan/tap
echo   • GET    /api/v1/reports/absensi
echo.
echo To stop: docker-compose down
echo To view logs: docker-compose logs -f backend
echo.
