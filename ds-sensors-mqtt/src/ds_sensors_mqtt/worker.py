import logging
import os
import time


class Worker():

    def __init__(self, sensor=None, publisher=None):
        if sensor is None:
            raise Exception('sensor is required')
        self._sensor = sensor

        if publisher is None:
            raise Exception('publisher is required')
        self._publisher = publisher

        self._interval = int(os.environ.get('WORKER_INTERVAL', '15'))

    def do(self):
        while True:
            try:
                value = self._sensor.do()
                self._publisher.do(value)
            except Exception as e:
                logging.exception(e)
            finally:
                time.sleep(self._interval)
                if os.environ.get('WORKER_RUN_ONCE') in ['true', 'True', '1']:
                    logging.info('once run, exiting')
                    break
