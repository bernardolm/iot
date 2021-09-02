import logging
import os
import time


class Worker():

    def __init__(self, home_assistant=None):
        if home_assistant is None:
            raise Exception('home_assistant is required')

        self._home_assistant = home_assistant
        self._interval = int(os.environ.get('WORKER_INTERVAL', '15'))

    def do(self):
        while True:
            try:
                self._home_assistant.do()
            except Exception as e:
                logging.exception(e)
            finally:
                if os.environ.get('WORKER_RUN_ONCE') in ['true', 'True', '1']:
                    logging.info('once run, exiting')
                    break
                time.sleep(self._interval)
