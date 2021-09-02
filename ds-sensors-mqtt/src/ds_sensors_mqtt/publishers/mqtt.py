import logging
import os
import socket
import time

from paho.mqtt.client import Client


class MQTTPublisher():

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

    def _log_success_sent(self, topic, payload):
        logging.info(f'sent {payload} to topic {topic}')

    def _log_failed_sent(self, topic, payload):
        logging.error(f'failed to send {payload} to topic {topic}')

    def publish(self, topic, payload, retain):
        try:
            result = self._client.publish(
                topic=topic,
                payload=payload,
                retain=retain)

            if result[0] == 0:
                self._log_success_sent(topic, payload)
            else:
                self._log_failed_sent(topic, payload)
                time.sleep(5)
                self.connect()
                self.publish(topic, payload, retain)

        except Exception as e:
            logging.exception(e)
            time.sleep(5)
            self.connect()
            self.publish(topic, payload, retain)

    def connect(self):
        self._client.connect(self._host, self._port)
