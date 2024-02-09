
KETOS_URL="https://github.com/Doer-org/ketos/releases/latest/download/ketos_Linux_arm64.tar.gz "

TEMP_DIR=$(mktemp -d)
echo "Created temporary directory: $TEMP_DIR"

cd $TEMP_DIR

echo "Downloading Ketos..."
curl -L $KETOS_URL -o ketos_linux_amd64.tar.gz

echo "Extracting files..."
tar -zxvf ketos_linux_amd64.tar.gz

echo "Installing Ketos to /usr/local/bin..."
sudo mv ketos /usr/local/bin/ketos

cd -
rm -rf $TEMP_DIR
echo "Removed temporary directory: $TEMP_DIR"

echo "Ketos installation is complete"
ketos --h
