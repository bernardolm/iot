#!/bin/sh

if test -f /etc/systemd/system/sensors-publisher-go.service; then
    eval "$SUDO" systemctl stop sensors-publisher-go
fi

eval "$SUDO" mv "${CURRENT_DIR}/service.debian" /etc/systemd/system/sensors-publisher-go.service

eval "$SUDO" systemctl daemon-reload
eval "$SUDO" systemctl enable sensors-publisher-go.service
eval "$SUDO" systemctl start sensors-publisher-go.service
eval "$SUDO" systemctl status sensors-publisher-go.service
