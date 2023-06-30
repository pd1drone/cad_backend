#!/bin/bash

# Create the file cad_backend.service with the specified content
cat << EOF > /etc/systemd/system/cad_app.service
[Unit]
Description=Raspberry CAD prototype 
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/
ExecStart=/usr/bin/python3 /root/cad_app.py
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the cad_backend.service file so it starts on boot
sudo systemctl enable cad_app.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the cad_backend.service
sudo systemctl start cad_app.service
