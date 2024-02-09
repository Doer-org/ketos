#!/bin/bash
KETOS_URL="https://github.com/Doer-org/ketos/releases/latest/download/ketos_Darwin_x86_64.tar.gz"

DOWNLOAD_DIR="$HOME/Downloads"

DOWNLOAD_PATH="$DOWNLOAD_DIR/ketos_Darwin_x86_64.tar.gz"

echo "Downloading Ketos..."
curl -L $KETOS_URL -o $DOWNLOAD_PATH

EXTRACT_PATH="$HOME/ketos"

if [ ! -d "$EXTRACT_PATH" ]; then
    mkdir -p "$EXTRACT_PATH"
fi

echo "Extracting files..."
tar -xzf $DOWNLOAD_PATH -C $EXTRACT_PATH

echo "Ketos has been installed to $EXTRACT_PATH"

echo "Please add $EXTRACT_PATH to your PATH to use Ketos from anywhere. You can do this by adding the following line to your ~/.bash_profile or ~/.zshrc:"
echo 'export PATH="$PATH:'$EXTRACT_PATH'"'
