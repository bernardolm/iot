import logging


class Publisher():

    def init(self, sensor=None):
        pass

    def do(self, value=None):
        if value is None:
            logging.info('publisher do!')
            return

        logging.info(f'publisher do with {value}!')

        # app.mqtt_publish(
        #     topic=f"homeassistant/{sensor.id}",
        #     payload=temperature_in_celsius,
        #     qos=1,
        #     retain=True
        # )
