#!/bin/bash

set -e

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

# Set installation directories based on OS
if [ "$OS" = "windows" ]; then
    INSTALL_DIR="$USERPROFILE/.basenuxt"
    BIN_DIR="$USERPROFILE/bin"
    BINARY_NAME="basenuxt.exe"
else
    INSTALL_DIR="$HOME/.basenuxt"
    BIN_DIR="/usr/local/bin"
    BINARY_NAME="basenuxt"
fi

# Create installation directories
mkdir -p "$INSTALL_DIR"
mkdir -p "$BIN_DIR" 2>/dev/null || {
    echo "Error: Unable to create $BIN_DIR directory. Please run with sudo:"
    echo "curl -sSL https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.sh | sudo bash"
    exit 1
}

echo "Installing BaseNuxt CLI..."
echo "OS: $OS"
echo "Architecture: $ARCH"

# Get the latest release version
LATEST_RELEASE=$(curl -s https://api.github.com/repos/BaseTechStack/basenuxt/releases/latest | grep "tag_name" | cut -d '"' -f 4)
if [ -z "$LATEST_RELEASE" ]; then
    echo "Warning: Could not determine latest version, using default v0.1.0"
    LATEST_RELEASE="v0.1.0"
fi

echo "Latest version: $LATEST_RELEASE"

# Download the appropriate binary
DOWNLOAD_URL="https://github.com/BaseTechStack/basenuxt/releases/download/$LATEST_RELEASE/basenuxt_${OS}_${ARCH}.tar.gz"
if [ "$OS" = "windows" ]; then
    DOWNLOAD_URL="https://github.com/BaseTechStack/basenuxt/releases/download/$LATEST_RELEASE/basenuxt_${OS}_${ARCH}.zip"
fi

echo "Downloading from: $DOWNLOAD_URL"
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

if [ "$OS" = "windows" ]; then
    curl -sL "$DOWNLOAD_URL" -o basenuxt.zip
    unzip basenuxt.zip
else
    curl -sL "$DOWNLOAD_URL" | tar xz
fi

# Install the binary
mv "$BINARY_NAME" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Create symlink
if [ "$OS" = "windows" ]; then
    # On Windows, copy to bin directory
    cp "$INSTALL_DIR/$BINARY_NAME" "$BIN_DIR/"
else
    # On Unix systems, create symlink with sudo
    echo "Creating symlink in $BIN_DIR (requires sudo)..."
    if ! sudo ln -sf "$INSTALL_DIR/$BINARY_NAME" "$BIN_DIR/$BINARY_NAME"; then
        echo "Error: Failed to create symlink. Please run the install script with sudo:"
        echo "curl -sSL https://raw.githubusercontent.com/BaseTechStack/basenuxt/main/install.sh | sudo bash"
        exit 1
    fi
fi

# Cleanup
cd - > /dev/null
rm -rf "$TMP_DIR"

echo "BaseNuxt CLI has been installed successfully!"
echo "Run 'basenuxt --help' to get started"

# Add to PATH for Windows if needed
if [ "$OS" = "windows" ]; then
    if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
        echo "Please add $BIN_DIR to your PATH to use the 'basenuxt' command"
        echo "You can do this by running:"
        echo "    setx PATH \"%PATH%;$BIN_DIR\""
    fi
fi
