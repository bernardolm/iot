import random


class Measurer():

    def __init__(self, sensor=None):
        self.name = f'measurer_test_{random.randint(1000, 9999)}'
        pass

    def value(self):
        return random.randint(12, 35) + (random.randint(100, 999)/100)
