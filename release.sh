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

# Check if tag already exists and delete it locally if it does
if git tag -l "$VERSION" | grep -q "$VERSION"; then
    echo "Tag $VERSION already exists locally, deleting it..."
    git tag -d "$VERSION"
fi

# Create tag locally
git tag -a "$VERSION" -m "Release $VERSION"

# Try to push tag to GitHub, but continue if it fails
echo "Attempting to push tag to GitHub (this may fail if you don't have push access)..."
if ! git push origin "$VERSION" --force; then
    echo "Warning: Failed to push tag to GitHub. Continuing with local build..."
fi

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

# Create GitHub release if gh command is available
RELEASE_NOTES="BaseNuxt CLI $VERSION

What's new:
$(git log --pretty=format:'- %s' $(git describe --tags --abbrev=0 HEAD^)..HEAD)

To upgrade BaseNuxt CLI, use:
\`\`\`bash
basenuxt upgrade
\`\`\`"

GITHUB_RELEASE_SUCCESS="false"
if command -v gh &> /dev/null; then
    echo "Attempting to create GitHub release (this may fail if you don't have proper permissions)..."
    if gh release create "$VERSION" \
        --title "BaseNuxt CLI $VERSION" \
        --notes "$RELEASE_NOTES" \
        basenuxt_*.{tar.gz,zip}; then
        GITHUB_RELEASE_SUCCESS="true"
        echo "GitHub release created successfully."
    else
        echo "Warning: Failed to create GitHub release. Archives are still available locally."
    fi
else
    echo "GitHub CLI (gh) not found. Skipping GitHub release creation."
    echo "The build artifacts are still available locally: basenuxt_*.{tar.gz,zip}"
fi

# Cleanup (only if GitHub release was successful)
if [ "$GITHUB_RELEASE_SUCCESS" = "true" ]; then
    echo "Cleaning up..."
    rm -f basenuxt_*.tar.gz basenuxt_*.zip
else
    echo "Keeping build artifacts for local use."
    echo "Files available in the current directory:"
    ls -la basenuxt_*.tar.gz basenuxt_*.zip 2>/dev/null || echo "No build artifacts found."
fi

echo "Release $VERSION completed successfully!"
