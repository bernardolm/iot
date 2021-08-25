import logging
import os

from dotenv import load_dotenv
from w1thermsensor import W1ThermSensor

from mqtt import MqttClient
from publisher import Publisher
from sensor import Sensor
from worker import Worker


def main():

    load_dotenv('config.env')

    log_level = os.environ.get('LOG_LEVEL', 'INFO').upper()
    logging.basicConfig(level=log_level)

    mc = MqttClient()

    for sensor in W1ThermSensor.get_available_sensors():
        logging.debug(sensor)

        try:
            s = Sensor(sensor)
            logging.debug(s)
            p = Publisher(sensor_name=s.name(), mqtt_client=mc)
            logging.debug(p)
            w = Worker(sensor=s, publisher=p)
            logging.debug(w)

            w.do()
        except Exception as e:
            logging.exception(sensor.id, e)


if __name__ == "__main__":
    main()
