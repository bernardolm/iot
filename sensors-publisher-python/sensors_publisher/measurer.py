import logging
import time


class Measurer:

    def __init__(self, sensor=None):
        if sensor is None:
            raise Exception('sensor is required')

        self._sensor = sensor
        self.name = self._sensor.id

    def value(self):
        try:
            temperature_in_celsius = self._sensor.get_temperature()
            logging.info(
                f'sensor {self._sensor.id} has temperature {temperature_in_celsius:.4f}')
            return temperature_in_celsius
        except Exception as e:
            logging.exception(
                ['sensor yet unavailable, trying again soon...', e])
            time.sleep(5)
            return self.value()
