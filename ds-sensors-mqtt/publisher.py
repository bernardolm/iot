import json
import logging
import os
import random

from paho.mqtt import client as mqtt_client

# NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/


class Publisher():

    def __init__(self, sensor_name=None):
        if sensor_name is None:
            raise Exception('sensor name is required')

        client_id = os.environ.get('MQTT_CLIENT_ID', sensor_name)
        self._client = mqtt_client.Client(f'{sensor_name}_client_id')

        user = os.environ.get('MQTT_USER')
        password = os.environ.get('MQTT_PASSWORD')
        if user is not None and password is not None:
            self._client.username_pw_set(
                username=user,
                password=password)

        self._host = os.environ.get('MQTT_HOST', 'localhost')
        self._port = int(os.environ.get('MQTT_PORT', '1883'))
        self._client.connect(self._host, self._port)

        self._state_topic = f'homeassistant/sensor/{sensor_name}/temperature/state'
        self._config_topic = f'homeassistant/sensor/{sensor_name}/temperature/config'

        rand = random.randint(1000, 9999)
        self._config_message = json.dumps({
            'device': {
                'identifiers': [
                    f'{sensor_name}_identifier',
                ],
                'model': 'DS18B20',
                'name': f'{sensor_name}_device_name',
                'unique_id': f'{sensor_name}_{rand}_unique_id',
            },
            'device_class': 'temperature',
            'friendly_name': f'Temperature from DS18B20 ({rand})',
            'name': sensor_name,
            'state_class': 'measurement',
            'state_topic': self._state_topic,
            'unique_id': f'{sensor_name}_{rand}_unique_id',
            'unit_of_measurement': '°C',
            'value_template': '{{ value_json.temperature }}',
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
