import logging
import random

from paho.mqtt import client as mqtt_client


class Publisher():

    def __init__(self, sensor=None):
        broker = '192.168.15.7'
        port = 1883
        self._topic = "/python/mqtt"
        client_id = 'ds-sensors-ha-app-daemon'
        self._client = mqtt_client.Client(client_id)
        self._client.connect(broker, port)

    def do(self, value=None):
        if value is None:
            logging.info('publisher has nothing to do...!')
            return

        logging.info(f'publishing {value}!')

        result = self._client.publish(self._topic, value)
        status = result[0]
        if status == 0:
            logging.debug(f'sent {value} to topic {self._topic}')
        else:
            logging.debug(f'failed to send {value} to topic {self._topic}')
