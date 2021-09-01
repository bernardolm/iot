import logging
import os
import socket

from paho.mqtt.client import Client


class MqttClient():

    def __init__(self):
        client_id = os.environ.get('MQTT_CLIENT_ID', socket.gethostname())
        self._client = Client(client_id)
        self._client.enable_logger(logger=logging)

        user = os.environ.get('MQTT_USER')
        password = os.environ.get('MQTT_PASSWORD')
        if user is not None and password is not None:
            self._client.username_pw_set(
                username=user,
                password=password)

        self._host = os.environ.get('MQTT_HOST', 'localhost')
        self._port = int(os.environ.get('MQTT_PORT', '1883'))
        self.connect()

    def publish(self, topic, payload, retain):
        result = self._client.publish(
            topic=topic,
            payload=payload,
            retain=retain)
        return result[0]

    def connect(self):
        self._client.connect(self._host, self._port)
