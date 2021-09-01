import random


class Sensor():

    def __init__(self, sensor=None):
        pass

    def name(self):
        return f'test_{random.randint(1000, 9999)}'

    def do(self):
        return random.randint(12, 35) + (random.randint(100, 999)/100)


class MqttClient():

    def publish(self, topic, payload, retain):
        return 0


class Sensors():

    def __init__(self):
        self.id = 1234

    def list(self):
        return [self]
