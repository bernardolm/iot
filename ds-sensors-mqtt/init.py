import logging

from w1thermsensor import W1ThermSensor

from publisher import Publisher
from sensor import Sensor
from worker import Worker

logging.basicConfig(level=logging.DEBUG)


def main():

    for sensor in W1ThermSensor.get_available_sensors():
        logging.debug(sensor)

        try:
            s = Sensor(sensor)
            logging.debug(s)
            p = Publisher(id=s.id())
            logging.debug(p)
            w = Worker(sensor=s, publisher=p)
            logging.debug(w)

            w.do()
        except Exception as e:
            logging.exception(sensor.id, e)


if __name__ == "__main__":
    main()
