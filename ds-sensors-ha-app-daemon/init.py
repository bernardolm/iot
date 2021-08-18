import logging

from w1thermsensor import W1ThermSensor

from .publisher import Publisher
from .sensor import Sensor
from .worker import Worker


def main():
    p = Publisher()
    logging.debug(p)

    for sensor in W1ThermSensor.get_available_sensors():
        logging.debug(sensor)

        s = Sensor(sensor)
        w = Worker(sensor=s, publisher=p)

        w.do()


if __name__ == "__main__":
    main()
