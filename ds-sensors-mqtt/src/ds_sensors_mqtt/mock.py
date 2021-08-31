
class Sensor():

    def __init__(self, sensor=None):
        pass

    def name(self):
        return 'DS18A20_test'

    def do(self):
        return 23.4567


class MqttClient():

    def publish(self, topic, payload, retain):
        return 0


class Sensors():

    def __init__(self):
        self.id = 1234
        pass

    def list(self):
        return [self]
