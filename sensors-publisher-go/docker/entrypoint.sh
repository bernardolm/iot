#!/bin/bash
# shellcheck shell=dash

# set -e

# rc-update show -v
# rc-status --servicelist
# rc-service sshd zap

rc-service rsyslog start
rc-service rsyslog status

rc-service sshd start
rc-service sshd status

sleep 9999
# /bin/bash
