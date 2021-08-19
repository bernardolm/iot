import logging


class Sensor():

    def __init__(self, sensor=None):
        if sensor is None:
            raise Exception('sensor is required')
        self._sensor = sensor

        try:
            self._sensor.set_resolution(9, persist=True)
        except Exception as e:
            logging.warning([self._sensor.id, e])

    def name(self):
        return f'DS18B20_{self._sensor.id}'

    def do(self):
        temperature_in_celsius = self._sensor.get_temperature()
        logging.debug(
            f'Sensor {self._sensor.id} has temperature {temperature_in_celsius:.4f}')

        return temperature_in_celsius
