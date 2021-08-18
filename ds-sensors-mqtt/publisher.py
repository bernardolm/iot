import logging
import random
import json

from paho.mqtt import client as mqtt_client

# NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/


class Publisher():

    def __init__(self, id=None):
        if id is None:
            id = andom.random()

        client_id = os.environ.get(
            'MQTT_CLIENT_ID', f'ds-sensors-mqtt_{id}').upper()
        self._client = mqtt_client.Client(client_id)

        self._host = os.environ.get('MQTT_HOST', 'localhost').upper()
        self._port = os.environ.get('MQTT_PORT', 1883).upper()
        self._client.connect(self._host, self._port)

        self._state_topic = f'homeassistant/sensor/{id}/state'
        self._config_topic = f'homeassistant/sensor/{id}/config'
        self._config_message = json.dumps({
            'name': id,
            'device_class': 'temperature',
            'state_topic': self._state_topic,
            'unit_of_measurement': '°C',
            'value_template': '{{value_json.temperature}}'
        })
        self._config()

    def _log_success_sent(self, topic, value):
        logging.info(f'sent {value} to topic {topic}')

    def _log_failed_sent(self, topic, value):
        logging.error(f'failed to send {value} to topic {topic}')

    def _publish(self, topic, value):
        result = self._client.publish(topic, value)
        status = result[0]
        if status == 0:
            self._log_success_sent(topic, value)
        else:
            self._log_failed_sent(topic, value)

    def _config(self):
        self._publish(self._config_topic, self._config_message)

    def _state(self, value):
        logging.debug(['value is', value])
        msg = json.dumps({'temperature': value})
        logging.debug(['message is', msg])
        self._publish(self._state_topic, msg)

    def do(self, value=None):
        if value is None:
            logging.info('publisher has nothing to do...!')
            return

        logging.info(f'publishing {value}!')

        self._state(value)
