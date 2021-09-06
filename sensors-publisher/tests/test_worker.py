import logging
import os
import unittest

from sensors_publisher.home_assistant import HomeAssistant
from sensors_publisher.measurer import Measurer
from sensors_publisher.publisher.mqtt import MQTTPublisher
from sensors_publisher.worker import Worker
from testfixtures import LogCapture


class TestWorker(unittest.TestCase):

    def test_worker(self):

        os.environ['WORKER_INTERVAL'] = '0'
        os.environ['WORKER_RUN_ONCE'] = 'true'

        p = MQTTPublisher()

        m = Measurer()

        ha = HomeAssistant(meansurer=m, publisher=p)

        w = Worker(home_assistant=ha)

        with LogCapture() as logs:
            await w.do()

        print(str(logs))

        assert 'publishing 23.4567' in str(logs)
        assert '[\'value is\', 23.4567]' in str(logs)
        assert '[\'message is\', \'{"temperature": 23.4567}\']' in str(logs)
        assert 'sent {"temperature": 23.4567} to topic homeassistant/sensor/' \
            'DS18A20_test/temperature/state' in str(logs)
