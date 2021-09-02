import json
import logging
import os

# NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/


class Publisher():

    def __init__(self, meansurer=None, mqtt_client=None):
        if meansurer is None:
            raise Exception('meansurer is required')

        self._availability_topic = f'ds18b20/bridge/state'
        self._config_topic = f'homeassistant/sensor/{meansurer.name}/temperature/config'
        self._meansurer = meansurer
        self._mqtt_client = mqtt_client
        self._state_topic = f'ds18b20/{meansurer.name}'

        self._availability()
        self._config()

    def _publish(self, topic, payload):
        if self._mqtt_client is None:
            logging.warning(f'mqtt client not configured, send only to stdout')
            status = 0
        else:
            status = self._mqtt_client.publish(
                topic=topic,
                payload=payload,
                retain=True)

    def _config(self):
        # NOTE: Ref.: https://www.home-assistant.io/docs/mqtt/discovery/
        id = f'ds18b20_temperature_sensor_{self._meansurer.name}'
        name = 'ds18b20 temperature sensor'
        payload = json.dumps({
            'availability': [
                {
                    'topic': self._availability_topic
                }
            ],
            'device': {
                'identifiers': [
                    id,
                ],
                'manufacturer': 'Unknown',
                'model': 'ds18b20',
                'name': name,
                'sw_version': 'ds18b20 0.0.1'
            },
            'device_class': 'temperature',
            'json_attributes_topic': self._state_topic,
            'name': f'{name} ({self._meansurer.name})',
            'state_class': 'measurement',
            'state_topic': self._state_topic,
            'unique_id': id,
            'unit_of_measurement': '°C',
            'value_template': '{{ value_json.temperature }}',
        })
        self._publish(self._config_topic, payload)

    def _availability(self):
        self._publish(self._availability_topic, 'online')

    def _state(self, value):
        payload = json.dumps({'temperature': value})
        self._publish(self._state_topic, payload)

    def do(self):
        value = self._meansurer.value()
        self._state(value)
