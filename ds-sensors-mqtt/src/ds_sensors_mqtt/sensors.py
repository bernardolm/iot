from w1thermsensor import W1ThermSensor


class Sensors():

    def list(self):
        return W1ThermSensor.get_available_sensors()
