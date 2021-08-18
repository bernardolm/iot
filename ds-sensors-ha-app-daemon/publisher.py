import logging


class Publisher():

    def __init__(self, sensor=None):
        pass

    def do(self, value=None):
        if value is None:
            logging.info('publisher has nothing to do...!')
            return

        logging.info(f'publisher do with {value}!')

        # app.mqtt_publish(
        #     topic=f"homeassistant/{sensor.id}",
        #     payload=temperature_in_celsius,
        #     qos=1,
        #     retain=True
        # )
