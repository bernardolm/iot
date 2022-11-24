#!/bin/bash

echo -e "starting service install"

mv /usr/local/sensor-publisher-go/sensor-publisher-go.service /etc/systemd/system/sensor-publisher-go.service

systemctl daemon-reload

/boot/dietpi/dietpi-services enable sensor-publisher-go
/boot/dietpi/dietpi-services start sensor-publisher-go
/boot/dietpi/dietpi-services status sensor-publisher-go

echo -e "finish service install"
