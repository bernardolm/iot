import os

from paho.mqtt import client as mqtt_client


class MqttClient():

    def __init__(self, sensor_name=None):
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

    def publish(self, topic, payload):
        result = self._client.publish(
            topic=topic,
            payload=payload,
            retain=True)
        return result[0]
