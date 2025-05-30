# Requires -RunAsAdministrator

# Function to get the latest release version
function Get-LatestRelease {
    $releaseUrl = "https://api.github.com/repos/BaseTechStack/bux/releases/latest"
    try {
        $release = Invoke-RestMethod -Uri $releaseUrl -ErrorAction Stop
        return $release.tag_name
    }
    catch {
        Write-Warning "Failed to get latest release: $_. Using default v0.1.0"
        return "v0.1.0"
    }
}

# Function to create directory if it doesn't exist
function Ensure-Directory {
    param([string]$Path)
    if (-not (Test-Path $Path)) {
        New-Item -ItemType Directory -Path $Path -Force | Out-Null
    }
}

Write-Host "Installing BUX CLI..." -ForegroundColor Green

# Detect architecture
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
Write-Host "Architecture: windows_$arch"

# Set installation paths
$installDir = Join-Path $env:USERPROFILE ".bux"
$binDir = Join-Path $env:USERPROFILE "bin"

# Create directories
Ensure-Directory $installDir
Ensure-Directory $binDir

# Get latest release
$version = Get-LatestRelease
Write-Host "Latest version: $version"

# Download URL
$downloadUrl = "https://github.com/BaseTechStack/bux/releases/download/$version/bux_windows_$arch.zip"
$zipPath = Join-Path $env:TEMP "bux.zip"
$exePath = Join-Path $installDir "bux.exe"
$binPath = Join-Path $binDir "bux.exe"

Write-Host "Downloading from: $downloadUrl"

try {
    # Download the zip file
    Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath -ErrorAction Stop

    # Extract the zip
    Expand-Archive -Path $zipPath -DestinationPath $installDir -Force

    # Copy to bin directory
    Copy-Item -Path $exePath -Destination $binPath -Force

    # Add to PATH if not already there
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($userPath -notlike "*$binDir*") {
        $newPath = "$userPath;$binDir"
        [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
        Write-Host "Added $binDir to PATH"
    }

    Write-Host "`nBUX CLI has been installed successfully!" -ForegroundColor Green
    Write-Host "Please restart your terminal to use the 'bux' command"
}
catch {
    Write-Error "Installation failed: $_"
    exit 1
}
finally {
    # Cleanup
    if (Test-Path $zipPath) {
        Remove-Item $zipPath -Force
    }
}

Write-Host "`nTo get started, run: bux --help"
