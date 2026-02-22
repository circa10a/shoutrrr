---
hide:
  - navigation
---
# Installation

## Overview

Shoutrrr is available via source code, released binaries, container images, and a GitHub Action.

## Build and Install from Source

```bash title="Install from Source"
go install github.com/circa10a/shoutrrr/shoutrrr@latest
```

## GitHub Release Binary Installation

The following scripts install the latest release binary to the user's `$HOME/go/bin` directory. Ensure this directory is in your `PATH` to enable execution without specifying the full binary path.
<!-- markdownlint-disable -->
=== "Windows (amd64)"

    ```powershell title="Windows (amd64) Installation"
    New-Item -ItemType Directory -Path $HOME\go\bin -Force | Out-Null; iwr (iwr https://api.github.com/repos/circa10a/shoutrrr/releases/latest | ConvertFrom-Json).assets.where({$_.name -like "*windows_amd64*.zip"}).browser_download_url -OutFile shoutrrr.zip; Add-Type -AssemblyName System.IO.Compression.FileSystem; ($z=[System.IO.Compression.ZipFile]::OpenRead("$PWD\shoutrrr.zip")).Entries | ? {$_.Name -eq 'shoutrrr.exe'} | % {[System.IO.Compression.ZipFileExtensions]::ExtractToFile($_, "$HOME\go\bin\$($_.Name)", $true)}; $z.Dispose(); rm shoutrrr.zip; if (Test-Path "$HOME\go\bin\shoutrrr.exe") { Write-Host "Successfully installed shoutrrr.exe to $HOME\go\bin" } else { Write-Host "Failed to install shoutrrr.exe" }
    ```

=== "Linux (amd64)"

    ```bash title="Linux (amd64) Installation"
    mkdir -p $HOME/go/bin && curl -L $(curl -s https://api.github.com/repos/circa10a/shoutrrr/releases/latest | grep -o 'https://[^"]*linux_amd64[^"]*\.tar\.gz') | tar -xz --strip-components=1 -C $HOME/go/bin shoutrrr
    ```

=== "macOS (amd64)"

    ```bash title="macOS (amd64) Installation"
    mkdir -p $HOME/go/bin && curl -L $(curl -s https://api.github.com/repos/circa10a/shoutrrr/releases/latest | grep -o 'https://[^"]*darwin_amd64[^"]*\.tar\.gz') | tar -xz --strip-components=1 -C $HOME/go/bin shoutrrr
    ```
<!-- markdownlint-restore -->
!!! Note
    Review the [release page](https://github.com/circa10a/shoutrrr/releases) for additional architectures (e.g., arm, arm64, i386, riscv64).

## Container Images

Shoutrrr provides lightweight Docker images based on Alpine Linux, supporting multiple architectures (amd64, arm64, arm/v6, i386, riscv64). Images are available on Docker Hub.

```bash title="Pull from Docker Hub"
docker pull circa10a/shoutrrr:latest
```

- **Repository**: <https://hub.docker.com/r/circa10a/shoutrrr>
- **Image Reference**: `circa10a/shoutrrr`
- **Tags**: `latest`, `vX.Y.Z` (e.g., `v0.8.0`), `latest-dev`, platform-specific (e.g., `amd64-latest`)

!!! Note
    Use `latest` for the latest stable release, versioned tags (e.g., `v0.8.0`) for specific releases, or `latest-dev` for development snapshots. Platform-specific tags are available for targeted deployments.

## Go Package

Add Shoutrrr to your Go project using:

```bash title="Add Go Package"
go get github.com/circa10a/shoutrrr@latest
```

## GitHub Action

Use Shoutrrr in GitHub workflows to send notifications.

```yaml title="Example GitHub Workflow with Shoutrrr"
name: Deploy
on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Some other steps needed for deploying
        run: ...
      - name: Shoutrrr
        uses: circa10a/shoutrrr-action@v1
        with:
          url: ${{ secrets.SHOUTRRR_URL }}
          title: Deployed ${{ github.sha }}
          message: See changes at ${{ github.event.compare }}.
```

!!! Note
    Pin the action to a specific SHA or version tag (e.g., `@v1`) and manage updates with Dependabot or Renovate for stability.
