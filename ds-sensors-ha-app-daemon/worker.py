import logging
import time


class Worker():

    def __init__(self, sensor=None, publisher=None):
        if sensor is None:
            raise Exception('sensor is required')
        self._sensor = sensor

        if publisher is None:
            raise Exception('publisher is required')
        self._publisher = publisher

    def do(self):

        while True:
            try:
                value = self._sensor.do()
                self._publisher.do(value)
            except Exception as e:
                logging.error(e)
            finally:
                time.sleep(5)
