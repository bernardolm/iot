from w1thermsensor import W1ThermSensor, Unit, Sensor
from pprint import pprint
import time

#for sensor in W1ThermSensor.get_available_sensors():
#    pprint(sensor)
#    pprint(sensor.__dict__)
#    #print("Sensor %s has temperature %.2f" % (sensor.id, sensor.get_temperature()))

sensor = W1ThermSensor(sensor_type=Sensor.DS18B20, sensor_id="020391776aaf")
pprint(sensor)
pprint(sensor.__dict__)
sensor.set_resolution(9, persist=True)

while True:
    temperature_in_celsius = sensor.get_temperature()
    print("Sensor %s has temperature %.2f" % (sensor.id, temperature_in_celsius))
    time.sleep(10)
