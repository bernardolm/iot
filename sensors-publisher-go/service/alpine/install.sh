#!/bin/sh

if [ "${USER}" != "root" ]; then
    export SUDO="sudo "
fi

${SUDO} rc-service sensors-publisher-go stop || true

${SUDO} rm -rf /var/log/sensors-publisher-go
${SUDO} mkdir -p /usr/share/sensors-publisher-go /var/log/sensors-publisher-go

${SUDO} mv /tmp/sensors-publisher-go/sensors-publisher-go /usr/share/sensors-publisher-go/sensors-publisher-go
${SUDO} mv /tmp/sensors-publisher-go/autostart /etc/init.d/sensors-publisher-go

${SUDO} chown root:root /etc/init.d/sensors-publisher-go

${SUDO} rc-update add sensors-publisher-go default
${SUDO} rc-service sensors-publisher-go start
${SUDO} rc-service sensors-publisher-go status
