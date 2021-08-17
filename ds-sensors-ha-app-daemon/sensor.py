import time
from pprint import pprint

from w1thermsensor import W1ThermSensor

import mqttapi as mqtt


class MyApp(mqtt.Mqtt):

    def initialize(self):
        pass


def get(app=None, sensor=None):
    if sensor is None:
        return

    temperature_in_celsius = sensor.get_temperature()
    print(f"Sensor {sensor.id} has temperature {temperature_in_celsius:.4f}")

    if app is None:
        return

    app.mqtt_publish(
        topic=f"homeassistant/{sensor.id}",
        payload=temperature_in_celsius,
        qos=1,
        retain=True
    )


def worker(app=None, sensor=None):
    if sensor is None:
        return

    try:
        sensor.set_resolution(9, persist=True)
    except Exception as e:
        print(sensor.id, e)

    while True:
        try:
            get(app, sensor)
        except Exception as e:
            print(e)
        finally:
            time.sleep(5)


def main():
    mqtt_cfg = mqtt.get_plugin_config()
    pprint(mqtt_cfg)
    pprint(mqtt_cfg.__dict__)

    app = MyApp()

    for sensor in W1ThermSensor.get_available_sensors():
        pprint(sensor)
        pprint(sensor.__dict__)
        worker(app, sensor)


if __name__ == "__main__":
    main()
