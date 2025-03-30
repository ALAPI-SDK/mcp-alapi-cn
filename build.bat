@echo off
setlocal

set VERSION=1.0.0
set BUILD_DIR=build

:: 创建构建目录
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

:: 清理旧的构建文件
echo Cleaning build directory...
del /Q %BUILD_DIR%\*

:: Windows 构建
echo Building for Windows/amd64...
set GOOS=windows
set GOARCH=amd64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-windows-amd64.exe" -ldflags="-s -w" main.go

echo Building for Windows/arm64...
set GOOS=windows
set GOARCH=arm64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-windows-arm64.exe" -ldflags="-s -w" main.go

:: Linux 构建
echo Building for Linux/amd64...
set GOOS=linux
set GOARCH=amd64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-linux-amd64" -ldflags="-s -w" main.go

echo Building for Linux/arm64...
set GOOS=linux
set GOARCH=arm64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-linux-arm64" -ldflags="-s -w" main.go

:: macOS 构建
echo Building for macOS/amd64...
set GOOS=darwin
set GOARCH=amd64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-darwin-amd64" -ldflags="-s -w" main.go

echo Building for macOS/arm64...
set GOOS=darwin
set GOARCH=arm64
go build -o "%BUILD_DIR%\mcp-alapi-cn-%VERSION%-darwin-arm64" -ldflags="-s -w" main.go

echo Build complete! Output files are in the build directory.
