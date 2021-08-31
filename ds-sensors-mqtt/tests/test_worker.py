import logging
import os
import unittest
from testfixtures import LogCapture
from src.ds_sensors_mqtt.publisher import Publisher
from src.ds_sensors_mqtt.worker import Worker


class SensorMock():

    def name(self):
        return 'DS18A20_test'

    def do(self):
        return 23.4567


class MqttClientMock():

    def publish(self, topic, payload, retain):
        return 0


class TestWorker(unittest.TestCase):

    def test_worker(self):

        os.environ['WORKER_INTERVAL'] = '0'
        os.environ['WORKER_RUN_ONCE'] = 'true'

        mc = MqttClientMock()

        s = SensorMock()

        p = Publisher(sensor_name=s.name(), mqtt_client=mc)

        w = Worker(sensor=s, publisher=p)

        with LogCapture() as logs:
            w.do()

        print(str(logs))

        assert 'publishing 23.4567' in str(logs)
        assert '[\'value is\', 23.4567]' in str(logs)
        assert '[\'message is\', \'{"temperature": 23.4567}\']' in str(logs)
        assert 'sent {"temperature": 23.4567} to topic homeassistant/sensor/' \
            'DS18A20_test/temperature/state' in str(logs)
