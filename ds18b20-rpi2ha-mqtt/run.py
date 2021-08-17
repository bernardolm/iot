import time
from pprint import pprint

from w1thermsensor import Sensor, W1ThermSensor


def get(sensor=None):
    if sensor is None:
        return

    temperature_in_celsius = sensor.get_temperature()
    print(f"Sensor {sensor.id} has temperature {temperature_in_celsius: %.4f}")


def main():
    sensor = W1ThermSensor(
        sensor_type=Sensor.DS18B20,
        sensor_id="020391776aaf")
    pprint(sensor)
    pprint(sensor.__dict__)

    try:
        sensor.set_resolution(9, persist=True)
    except Exception as e:
        print(e)

    while True:
        try:
            get(sensor)
        except Exception as e:
            print(e)
        finally:
            time.sleep(5)


if __name__ == "__main__":
    main()
