#!/bin/bash

# Create the file cad_backend.service with the specified content
cat << EOF > /etc/systemd/system/cad_backend.service
[Unit]
Description=Raspberry CAD backend http server
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/cad_backend/
ExecStart=/root/cad_backend/cmd/cad
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the cad_backend.service file so it starts on boot
sudo systemctl enable cad_backend.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the cad_backend.service
sudo systemctl start cad_backend.service
