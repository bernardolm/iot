from distutils.core import setup

from setuptools import setup

with open('requirements.txt') as f:
    reqs = f.read().strip().split('\n')
    reqs = [r for r in reqs if '-e' not in r]

setup(
    name="sensors_publisher",
    version="0.0.1",
    install_requires=reqs,
    packages=[
        "sensors_publisher.interface",
        "sensors_publisher.publisher",
        "sensors_publisher.sensor",
        "sensors_publisher",
    ]
)
