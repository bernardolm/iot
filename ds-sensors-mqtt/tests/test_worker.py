from src.ds_sensors_mqtt.publisher import Publisher
from src.ds_sensors_mqtt.worker import Worker


class SensorMock():

    def name(self):
        return 'DS18A20_test'

    def do(self):
        return 23.4567


class MqttClientMock():

    def publish(self, topic, payload, retain):
        return 0


def test_worker():

    mc = MqttClientMock()

    s = SensorMock()

    p = Publisher(sensor_name=s.name(), mqtt_client=mc)

    w = Worker(sensor=s, publisher=p)

    w.do()

    assert True == True
