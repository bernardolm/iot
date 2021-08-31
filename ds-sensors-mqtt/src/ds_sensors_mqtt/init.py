import logging
import multiprocessing
import os
import sys
import threading
import time

from dotenv import load_dotenv
from src.ds_sensors_mqtt.publisher import Publisher
from src.ds_sensors_mqtt.updater import Updater
from src.ds_sensors_mqtt.worker import Worker


def main():

    load_dotenv('config.env', override=True)
    log_level = os.environ.get('LOG_LEVEL', 'INFO').upper()
    logging.basicConfig(level=log_level)

    debug = os.environ.get('DEBUG') in ['true', 'True', '1']

    logging.info(f'running in {"debug" if debug else "normal"} mode...')

    if debug:
        from src.ds_sensors_mqtt.mock import MqttClient, Sensor, Sensors
    else:
        from src.ds_sensors_mqtt.mqtt import MqttClient
        from src.ds_sensors_mqtt.sensor import Sensor
        from src.ds_sensors_mqtt.sensors import Sensors

    jobs = []
    event = multiprocessing.Event()

    mc = MqttClient()

    for sensor in Sensors().list():
        logging.debug(sensor)

        try:
            s = Sensor(sensor)
            logging.debug(s)

            p = Publisher(sensor_name=s.name(), mqtt_client=mc)
            logging.debug(p)

            w = Worker(sensor=s, publisher=p)
            logging.debug(w)

            p = multiprocessing.Process(target=w.do)
            p.start()
            jobs.append(p)

        except Exception as e:
            logging.exception([sensor.id, e])

    u = Updater()
    p = multiprocessing.Process(target=u.do, args=(event,))
    p.start()
    jobs.append(p)

    while True:
        if event.is_set():
            logging.warning('event is set, terminating jobs...')
            for i in jobs:
                logging.warning(f'terminating job {i}')
                i.terminate()
            logging.warning('exiting')
            sys.exit()
        time.sleep(5)


if __name__ == "__main__":
    main()
