import logging
import random

from paho.mqtt import client as mqtt_client


class Publisher():

    def __init__(self, sensor=None):
        client_id = 'ds-sensors-ha-app-daemon'
        self._client = mqtt_client.Client(client_id)
        broker = '192.168.15.7'
        port = 1883
        self._client.connect(broker, port)

        self._state_topic = f'homeassistant/sensor/{sensor.id}/state'
        self._config_topic = f'homeassistant/sensor/{sensor.id}/config'
        self._config_message = f`{
            "name":"{sensor.id}",
            "device_class":"temperature",
            "state_topic":"{self._state_topic}",
            "unit_of_measurement": "°C",
            "value_template": "{{ value_json.temperature}}"
        }`
        self._config()

    def _log_success_sent(self, topic, value):
        logging.info(f'sent {value} to topic {topic}')

    def _log_failed_sent(self, topic, value):
        logging.error(f'failed to send {value} to topic {topic}')

    def _publish(self, topic, value):
        result = self._client.publish(topic, value)
        status = result[0]
        if status == 0:
            self._log_success_sent(value, topic)
        else:
            self._log_failed_sent(value, topic)

    def _config(self):
        self._publish(self._config_topic, self._config_message)

    def _state(self, value):
        msg = f'{"temperature":{value}}'
        self._publish(self._state_topic, msg)

    def do(self, value=None):
        if value is None:
            logging.info('publisher has nothing to do...!')
            return

        logging.info(f'publishing {value}!')

        try:
            self._state(value)
        except:
            raise e
