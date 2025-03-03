#!/bin/bash

set -e

# Check if version is provided
if [ -z "$1" ]; then
    echo "Usage: ./release.sh <version>"
    echo "Example: ./release.sh 1.0.0"
    exit 1
fi

VERSION="v$1"
BUILD_DATE=$(date -u '+%Y-%m-%dT%H:%M:%SZ')
GO_VERSION=$(go version | cut -d' ' -f3)
COMMIT_HASH=$(git rev-parse HEAD)
PLATFORMS=("darwin_amd64" "darwin_arm64" "linux_amd64" "windows_amd64")

echo "Creating release $VERSION..."

# Create and push tag
git tag -a "$VERSION" -m "Release $VERSION"
git push origin "$VERSION"

# Build binaries for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
    echo "Building for $PLATFORM..."
    
    # Split platform into OS and ARCH
    IFS='_' read -r OS ARCH <<< "$PLATFORM"
    
    # Set output binary name
    if [ "$OS" = "windows" ]; then
        BINARY="basenuxt.exe"
    else
        BINARY="basenuxt"
    fi

    # Build
    GOOS=$OS GOARCH=$ARCH go build \
        -ldflags "-X 'github.com/BaseTechStack/basenuxt/version.Version=$VERSION' \
                  -X 'github.com/BaseTechStack/basenuxt/version.CommitHash=$COMMIT_HASH' \
                  -X 'github.com/BaseTechStack/basenuxt/version.BuildDate=$BUILD_DATE' \
                  -X 'github.com/BaseTechStack/basenuxt/version.GoVersion=$GO_VERSION'" \
        -o "$BINARY"

    # Create archive
    if [ "$OS" = "windows" ]; then
        zip "basenuxt_${PLATFORM}.zip" "$BINARY"
        rm "$BINARY"
    else
        tar czf "basenuxt_${PLATFORM}.tar.gz" "$BINARY"
        rm "$BINARY"
    fi
done

# Create GitHub release
RELEASE_NOTES="BaseNuxt CLI $VERSION

What's new:
$(git log --pretty=format:'- %s' $(git describe --tags --abbrev=0 HEAD^)..HEAD)

To upgrade BaseNuxt CLI, use:
\`\`\`bash
basenuxt upgrade
\`\`\`"

gh release create "$VERSION" \
    --title "BaseNuxt CLI $VERSION" \
    --notes "$RELEASE_NOTES" \
    basenuxt_*.{tar.gz,zip}

# Cleanup
echo "Cleaning up..."
rm -f basenuxt_*.tar.gz basenuxt_*.zip

echo "Release $VERSION completed successfully!"
