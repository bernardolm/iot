FROM alpine:3.19

ARG SSH_PUB_KEY

WORKDIR /tmp

RUN apk update \
    && apk add --update --no-cache \
        bash coreutils nano openrc openssh-server rsync rsyslog util-linux

RUN mkdir -p /run/openrc /dev/log /var/log/ /etc/rsyslog.d \
        /etc/conf.d /etc/ssh/ssh_config.d /root/.ssh/ \
    && touch /run/openrc/softlevel \
    && sed -i \
        's/module(load="imuxsock")/# module(load="imuxsock")/g' \
        /etc/rsyslog.conf

COPY etc_rsyslog.d_custom.conf /etc/rsyslog.d/custom.conf
COPY etc_conf.d_rc.conf /etc/conf.d/rc.conf
COPY etc_ssh_ssh_config.d_custom.conf /etc/ssh/ssh_config.d/custom.conf

RUN rc-update add rsyslog default \
    && rc-update add sshd default \
    && rc-update --update

RUN ssh-keygen -A \
    && echo "$SSH_PUB_KEY" > /root/.ssh/authorized_keys

VOLUME ["/sys/fs/cgroup"]

EXPOSE 22

COPY entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]
