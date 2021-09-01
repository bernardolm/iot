from w1thermsensor import W1ThermSensor


class Sensors():

    def list(self):
        sensors = []

        for s in W1ThermSensor.get_available_sensors():
            try:
                s.set_resolution(9, persist=True)
            except Exception as e:
                logging.exception([s.id, e])

            sensors.append(s)

        return sensors
