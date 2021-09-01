from distutils.core import setup

from setuptools import setup

with open('requirements.txt') as f:
    reqs = f.read().strip().split('\n')
    reqs = [r for r in reqs if '-e' not in r]

setup(
    name="ds_sensors_mqtt",
    version="0.0.1",
    install_requires=reqs,
    packages=[
        "src.ds_sensors_mqtt",
        "src.ds_sensors_mqtt.mocks"
    ]
)
