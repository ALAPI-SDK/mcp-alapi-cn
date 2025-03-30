#!/bin/bash

# 版本号
VERSION="1.0.0"

# 创建构建目录
mkdir -p build

# 构建函数
build() {
    local GOOS=$1
    local GOARCH=$2
    local SUFFIX=$3
    
    echo "Building for $GOOS/$GOARCH..."
    
    # 设置输出文件名
    local OUTPUT="build/mcp-alapi-cn-${VERSION}-${GOOS}-${GOARCH}${SUFFIX}"
    
    # 执行构建
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT" -ldflags="-s -w" main.go
    
    # 检查构建结果
    if [ $? -eq 0 ]; then
        echo "✓ Successfully built $OUTPUT"
    else
        echo "✗ Failed to build for $GOOS/$GOARCH"
        exit 1
    fi
}

# 清理旧的构建文件
echo "Cleaning build directory..."
rm -rf build/*

# Windows 构建
build "windows" "amd64" ".exe"
build "windows" "arm64" ".exe"

# Linux 构建
build "linux" "amd64" ""
build "linux" "arm64" ""

# macOS 构建
build "darwin" "amd64" ""
build "darwin" "arm64" ""

echo "Build complete! Output files are in the build directory."
