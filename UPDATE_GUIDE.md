# BUX Update Guide

This document provides instructions for updating the BUX CLI to the latest version.

## Automatic Update

The BUX CLI includes a built-in update command that will automatically download and install the latest version:

```bash
bux update
```

This command will:
1. Check for the latest version on GitHub
2. Download the appropriate binary for your system
3. Replace the current binary with the new one

## Manual Update

If you prefer to update manually, you can follow these steps:

### macOS and Linux

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | bash
```

If you need to install in a protected directory (like `/usr/local/bin`), use:

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | sudo bash
```

### Windows

#### Option 1: Using PowerShell

```powershell
iwr -useb https://raw.githubusercontent.com/BaseTechStack/bux/main/install.ps1 | iex
```

#### Option 2: Using Git Bash

```bash
curl -sSL https://raw.githubusercontent.com/BaseTechStack/bux/main/install.sh | bash
```

## Verifying the Update

After updating, you can verify the installed version with:

```bash
bux --version
```

## Downgrading to a Specific Version

If you need to install a specific version of BUX CLI, you can download the binary directly from the GitHub releases page:

https://github.com/BaseTechStack/bux/releases

Download the appropriate archive for your system, extract it, and place the binary in your PATH.

## Troubleshooting

If you encounter issues during the update process:

1. **Permission Denied**: Make sure you have the necessary permissions to write to the installation directory. On Unix systems, you may need to use `sudo`.

2. **Binary Not Found**: Ensure that the installation directory is in your PATH.

3. **Version Not Updated**: Check if you have multiple copies of BUX CLI installed on your system. Use `which bux` (Unix) or `where bux` (Windows) to locate all instances.

4. **Network Issues**: If you're behind a proxy or firewall, you may need to configure your network settings to allow downloads from GitHub.

For further assistance, please open an issue on the [GitHub repository](https://github.com/BaseTechStack/bux/issues).
