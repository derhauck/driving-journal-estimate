#!/bin/bash

# Identify OS
OS="unknown"
case "$(uname -s)" in
    Darwin) OS="mac";;
    Linux) OS="linux";;
esac

# Fetch the latest release download URL for the correct OS binary using GitHub API
URL=$(curl -s https://api.github.com/repos/DERHauck/driving-journal-estimate/releases/latest |
grep "browser_download_url.*${OS}-bin.tar.gz" |
cut -d '"' -f 4)

# Define installation directory and binary path
INSTALL_DIR="/usr/local/bin"
BINARY_PATH="${INSTALL_DIR}/dje"

# Download and extract the binary
curl -L $URL | tar xz -C $INSTALL_DIR

# Make the binary executable
chmod +x $BINARY_PATH

echo "Installation complete. You can now use 'dje' from anywhere."
