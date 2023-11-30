#!/bin/bash
set -ex

REPO_NAME=$1
DOMAIN=$2


# Ensure the production directory exists
mkdir -p /home/production/"$REPO_NAME" &&
cd /home/production/"$REPO_NAME" &&

# Ensure the production directory exists
mkdir -p /home/production/"$REPO_NAME" &&
cd /home/production/"$REPO_NAME" &&

# If it's a Git repo, pull changes, otherwise clone the repository
if [ -d "$REPO_NAME/.git" ]; then
    git -C "$REPO_NAME" pull
else
    git clone git@github.com:cploutarchou/CryptoGainerAPI-Client.git "$REPO_NAME"
fi &&


# Change to the repository directory and pull the latest changes
cd "$REPO_NAME" &&
git pull &&

# Build the Go application
go build -o "$REPO_NAME" &&

# Remove the existing systemd service file if it exists
sudo rm -f /etc/systemd/system/"$REPO_NAME".service

# Create a new systemd service file for the application
sudo tee /etc/systemd/system/"$REPO_NAME".service > /dev/null << EOF
[Unit]
Description=$REPO_NAME Service
After=network.target

[Service]
Type=simple
User=root
ExecStart=/home/production/$REPO_NAME/$REPO_NAME/$REPO_NAME
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd to apply the new service file, enable and restart the service
sudo systemctl daemon-reload &&
sudo systemctl enable "$REPO_NAME" &&
sudo systemctl restart "$REPO_NAME" &&

# Set up Nginx configuration if it doesn't exist
if [ ! -f /etc/nginx/sites-available/"$REPO_NAME".conf ]; then
  sudo tee /etc/nginx/sites-available/"$REPO_NAME".conf > /dev/null <<EOF
  server {
    listen 80;
    server_name $DOMAIN;

    location / {
      proxy_pass http://127.0.0.1:8999;
      proxy_http_version 1.1;
      proxy_set_header Upgrade \$http_upgrade;
      proxy_set_header Connection 'upgrade';
      proxy_set_header Host \$host;
      proxy_cache_bypass \$http_upgrade;
    }
  }
EOF

  # Enable the Nginx site and reload the Nginx service
  sudo ln -s /etc/nginx/sites-available/"$REPO_NAME".conf /etc/nginx/sites-enabled/
  sudo systemctl reload nginx

  # Obtain an SSL certificate
  sudo certbot --nginx -d "$DOMAIN" --non-interactive --agree-tos --email cploutarchou@gmail.com
fi
