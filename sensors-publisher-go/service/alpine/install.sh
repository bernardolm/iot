#!/bin/sh

if [ "${USER}" != "root" ]; then
    export SUDO="sudo "
fi

INSTALL_PATH=/usr/share/sensors-publisher-go
LOG_PATH=/var/log/sensors-publisher-go
TMP_PATH=/tmp/sensors-publisher-go

${SUDO} rc-service sensors-publisher-go stop || true
${SUDO} rc-update delete sensors-publisher-go || true
${SUDO} killall -9 sensors-publisher-go || true

${SUDO} rm -rf "${LOG_PATH}/*" "${INSTALL_PATH}/*" || \
        mkdir -p "${LOG_PATH}" "${INSTALL_PATH}"

${SUDO} mv -f "${TMP_PATH}/sensors-publisher-go" "${INSTALL_PATH}/sensors-publisher-go"
${SUDO} mv -f "${TMP_PATH}/autostart" /etc/init.d/sensors-publisher-go

ls -lh "/etc/sensors-publisher-go"
ls -lh "${INSTALL_PATH}"
ls -lh "${LOG_PATH}"
ls -lh "${TMP_PATH}"

${SUDO} rc-update add sensors-publisher-go default
${SUDO} rc-service sensors-publisher-go start
${SUDO} rc-service sensors-publisher-go status

echo "installation completed successfully. now only showing recent log."
echo "you can exit with ctrl+c"

${SUDO} tail -n20 -f /var/log/sensors-publisher-go/stderr.log
