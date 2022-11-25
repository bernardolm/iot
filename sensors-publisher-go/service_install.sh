#!/bin/bash

echo -e "starting service install"

systemctl stop sensor-publisher-go

mv /usr/local/sensor-publisher-go/sensor-publisher-go.service /etc/systemd/system/sensor-publisher-go.service

systemctl daemon-reload

systemctl enable sensor-publisher-go.service
systemctl start sensor-publisher-go.service
systemctl status sensor-publisher-go.service

echo -e "finish service install"
