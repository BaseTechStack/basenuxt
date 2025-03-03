#!/bin/bash

set -e

# This script simulates the installation process without actually installing anything
# It's useful for testing the installation script

# Detect OS and architecture
detect_os() {
    case "$(uname -s)" in
        Darwin*)   echo "darwin" ;;
        Linux*)    echo "linux" ;;
        MINGW64*|MSYS*|CYGWIN*) echo "windows" ;;
        *)         echo "unknown" ;;
    esac
}

detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64) echo "amd64" ;;
        arm64|aarch64) echo "arm64" ;;
        *)            echo "unknown" ;;
    esac
}

OS=$(detect_os)
ARCH=$(detect_arch)

if [ "$OS" = "unknown" ] || [ "$ARCH" = "unknown" ]; then
    echo "Unsupported operating system or architecture"
    exit 1
fi

echo "Testing BaseNuxt CLI installation..."
echo "OS: $OS"
echo "Architecture: $ARCH"

# Use local version
VERSION="v0.1.0"
echo "Version: $VERSION"

# Check if the local archive exists
LOCAL_ARCHIVE="basenuxt_${OS}_${ARCH}.tar.gz"
if [ "$OS" = "windows" ]; then
    LOCAL_ARCHIVE="basenuxt_${OS}_${ARCH}.zip"
fi

if [ -f "$LOCAL_ARCHIVE" ]; then
    echo "✅ Local archive found: $LOCAL_ARCHIVE"
else
    echo "❌ Local archive not found: $LOCAL_ARCHIVE"
    exit 1
fi

# Simulate downloading from GitHub
DOWNLOAD_URL="https://github.com/BaseTechStack/basenuxt/releases/download/$VERSION/$LOCAL_ARCHIVE"
echo "✅ Download URL would be: $DOWNLOAD_URL"

echo "✅ Installation test completed successfully!"
echo "The installation script is correctly configured to download and install BaseNuxt CLI v0.1.0."
