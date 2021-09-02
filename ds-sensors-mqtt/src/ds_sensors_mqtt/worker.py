import logging
import os
import time


class Worker():

    def __init__(self, publisher=None):
        if publisher is None:
            raise Exception('publisher is required')

        self._publisher = publisher
        self._interval = int(os.environ.get('WORKER_INTERVAL', '15'))

    def do(self):
        while True:
            try:
                self._publisher.do()
            except Exception as e:
                logging.exception(e)
            finally:
                if os.environ.get('WORKER_RUN_ONCE') in ['true', 'True', '1']:
                    logging.info('once run, exiting')
                    break
                time.sleep(self._interval)
