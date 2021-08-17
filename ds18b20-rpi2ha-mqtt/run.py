from w1thermsensor import W1ThermSensor, Unit, Sensor
from pprint import pprint
import time

sensor = W1ThermSensor(sensor_type=Sensor.DS18B20, sensor_id="020391776aaf")
pprint(sensor)
pprint(sensor.__dict__)
sensor.set_resolution(persist=True)

def get():
    temperature_in_celsius = sensor.get_temperature()
    print("Sensor %s has temperature %.2f" % (sensor.id, temperature_in_celsius))

while True:
    try:
        get()
    except Exception as e:
        print(e)

    time.sleep(10)
