import logging
import os
import unittest

from src.ds_sensors_mqtt.mock import MqttClient, Sensor, Sensors
from src.ds_sensors_mqtt.publisher import Publisher
from src.ds_sensors_mqtt.worker import Worker
from testfixtures import LogCapture


class TestWorker(unittest.TestCase):

    def test_worker(self):

        os.environ['WORKER_INTERVAL'] = '0'
        os.environ['WORKER_RUN_ONCE'] = 'true'

        mc = MqttClient()

        s = Sensor()

        p = Publisher(sensor_name=s.name(), mqtt_client=mc)

        w = Worker(sensor=s, publisher=p)

        with LogCapture() as logs:
            await w.do()

        print(str(logs))

        assert 'publishing 23.4567' in str(logs)
        assert '[\'value is\', 23.4567]' in str(logs)
        assert '[\'message is\', \'{"temperature": 23.4567}\']' in str(logs)
        assert 'sent {"temperature": 23.4567} to topic homeassistant/sensor/' \
            'DS18A20_test/temperature/state' in str(logs)
