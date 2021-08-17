from w1thermsensor import W1ThermSensor, Unit

for sensor in W1ThermSensor.get_available_sensors():
    print("Sensor %s has temperature %.2f" % (sensor.id, sensor.get_temperature()))

sensor = W1ThermSensor(sensor_type=Sensor.DS18B20, sensor_id="28-020391776aaf")
sensor.set_resolution(9, persist=True)
temperature_in_celsius = sensor.get_temperature()
