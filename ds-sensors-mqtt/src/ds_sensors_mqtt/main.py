import logging
import multiprocessing
import os
import sys
import threading
import time

from dotenv import load_dotenv
from src.ds_sensors_mqtt.home_assistant import HomeAssistant
from src.ds_sensors_mqtt.updater import Updater
from src.ds_sensors_mqtt.worker import Worker
from src.ds_sensors_mqtt.publishers.mqtt import MQTTPublisher


def main():
    load_dotenv('config.env', override=True)
    log_level = os.environ.get('LOG_LEVEL', 'INFO').upper()
    logging.basicConfig(level=log_level)

    debug = os.environ.get('DEBUG') in ['true', 'True', '1']

    logging.info(f'running in {"debug" if debug else "normal"} mode...')

    if debug:
        from src.ds_sensors_mqtt.mocks.measurer import Measurer
        from src.ds_sensors_mqtt.mocks.sensors import Sensors
    else:
        from src.ds_sensors_mqtt.measurer import Measurer
        from src.ds_sensors_mqtt.sensors import Sensors

    jobs = []
    event = multiprocessing.Event()

    p = MQTTPublisher()

    for sensor in Sensors().list():
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
