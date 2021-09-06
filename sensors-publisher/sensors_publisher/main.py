import logging
import multiprocessing
import os
import sys
import threading
import time

from dotenv import load_dotenv
from sensors_publisher.home_assistant import HomeAssistant
from sensors_publisher.updater import Updater
from sensors_publisher.worker import Worker
from sensors_publisher.publisher.mqtt import MQTTPublisher


def main():
    load_dotenv('config.env', override=True)
    log_level = os.environ.get('LOG_LEVEL', 'INFO').upper()
    logging.basicConfig(level=log_level)

    debug = os.environ.get('DEBUG') in ['true', 'True', '1']

    logging.info(f'running in {"debug" if debug else "normal"} mode...')

    if debug:
        from sensors_publisher.mock.measurer import Measurer
        from sensors_publisher.mock.sensors import DS18A20
    else:
        from sensors_publisher.measurer import Measurer
        from sensors_publisher.sensors import DS18A20

    jobs = []
    event = multiprocessing.Event()

    p = MQTTPublisher()

    for sensor in DS18A20().list():
        logging.debug(sensor)

        try:
            m = Measurer(sensor)
            logging.debug(m)

            ha = HomeAssistant(meansurer=m, publisher=p)
            logging.debug(ha)

            w = Worker(home_assistant=ha)
            logging.debug(w)

            pc = multiprocessing.Process(target=w.do)
            pc.start()
            jobs.append(pc)

        except Exception as e:
            logging.exception(e)

    u = Updater()
    pc = multiprocessing.Process(target=u.do, args=(event,))
    pc.start()
    jobs.append(pc)

    while True:
        if event.is_set():
            logging.warning(
                'multiprocessing event is set, terminating jobs...')
            for i in jobs:
                logging.warning(f'terminating job with pid {i.pid}')
                i.terminate()
            logging.warning('exiting')
            sys.exit()
        time.sleep(5)


if __name__ == "__main__":
    main()
