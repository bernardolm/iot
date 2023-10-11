#!/bin/sh

eval "$SUDO service sensors-publisher-go stop" || true

eval "$SUDO" mkdir -p /usr/share/sensors-publisher-go /var/log/sensors-publisher-go
eval "$SUDO" mv "${CURRENT_DIR}/config.env" /usr/share/sensors-publisher-go/config.env
eval "$SUDO" mv "${CURRENT_DIR}/sensors-publisher-go" /usr/share/sensors-publisher-go/sensors-publisher-go

eval "$SUDO" mv "${CURRENT_DIR}/autostart.alpine" /etc/init.d/sensors-publisher-go

eval "$SUDO" chown root:root /etc/init.d/sensors-publisher-go
eval "$SUDO" chown root:root /usr/share/sensors-publisher-go/config.env

# eval "$SUDO" mkdir -p /var/log/sensors-publisher-go

eval "$SUDO" rc-update add sensors-publisher-go default
eval "$SUDO" service sensors-publisher-go start
eval "$SUDO" service sensors-publisher-go status
