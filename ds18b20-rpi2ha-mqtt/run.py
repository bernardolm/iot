from w1thermsensor import W1ThermSensor, Unit, Sensor
from pprint import pprint
import time

def get(sensor=None):
    if sensor is None:
        return

    temperature_in_celsius = sensor.get_temperature()
    print("Sensor %s has temperature %.4f" % (sensor.id, temperature_in_celsius))

def main():
    sensor = W1ThermSensor(sensor_type=Sensor.DS18B20, sensor_id="020391776aaf")
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
            time.sleep(10)

if __name__ == "__main__":
    main()
