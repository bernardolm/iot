import json
import logging
import os

# NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/


class Publisher():

    def __init__(self, sensor_name=None, mqtt_client=None):
        if sensor_name is None:
            raise Exception('sensor name is required')

        self._availability_topic = f'ds18b20/bridge/state'
        self._config_topic = f'homeassistant/sensor/{sensor_name}/temperature/config'
        self._mqtt_client = mqtt_client
        self._sensor_name = sensor_name
        self._state_topic = f'ds18b20/{sensor_name}'

        self._config()
        self._availability()

    def _log_success_sent(self, topic, payload):
        logging.info(f'sent {payload} to topic {topic}')

    def _log_failed_sent(self, topic, payload):
        logging.error(f'failed to send {payload} to topic {topic}')
        self._mqtt_client.connect()

    def _publish(self, topic, payload):
        if self._mqtt_client is None:
            logging.warning(f'mqtt client not configured, send only to stdout')
            status = 0
        else:
            status = self._mqtt_client.publish(
                topic=topic,
                payload=payload,
                retain=True)

        if status == 0:
            self._log_success_sent(topic, payload)
        else:
            self._log_failed_sent(topic, payload)

    def _config(self):
        # NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/
        payload = json.dumps({
            'availability': [
                {
                    'topic': self._availability_topic
                }
            ],
            'device': {
                'identifiers': [
                    f'ds18b20_{self._sensor_name}',
                ],
                'manufacturer': 'Unknown',
                'model': 'ds18b20',
                'name': self._sensor_name,
                'sw_version': 'ds18b20 0.0.1'
            },
            'device_class': 'temperature',
            'json_attributes_topic': self._state_topic,
            'name': self._sensor_name,
            'state_class': 'measurement',
            'state_topic': self._state_topic,
            'unique_id': f'{self._sensor_name}_temperature_ds18b20',
            'unit_of_measurement': '°C',
            'value_template': '{{ value_json.temperature }}',
        })
        self._publish(self._config_topic, payload)

    def _availability(self):
        self._publish(self._availability_topic, 'online')

    def _state(self, value):
        payload = json.dumps({'temperature': value})
        self._publish(self._state_topic, payload)

    def do(self, value=None):
        if value is None:
            logging.debug('publisher has nothing to do...!')
            return

        logging.debug(f'publishing {value}!')

        self._state(value)
