import logging
import time


class Sensor():

    def __init__(self, sensor=None):
        if sensor is None:
            raise Exception('sensor is required')
        self._sensor = sensor

        try:
            self._sensor.set_resolution(9, persist=True)
        except Exception as e:
            logging.exception([self._sensor.id, e])

    def name(self):
        self.do()
        return self._sensor.id

    def do(self):
        try:
            temperature_in_celsius = self._sensor.get_temperature()
            logging.info(
                f'sensor {self._sensor.id} has temperature {temperature_in_celsius:.4f}')
        except Exception as e:
            logging.exception(
                ['sensor yet unavailable, trying again soon...', e])
            time.sleep(3)
            return self.do()

        return temperature_in_celsius
