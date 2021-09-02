ENV?=development
ZOMBIES=$(shell ps aux | grep src/ds_sensors_mqtt/main.py | grep -v grep | awk '{print $$2}')


install:
ifeq ("${ENV}","production")
	@echo 'installing app...'
	@curl -sL https://github.com/bernardolm/iot/archive/refs/heads/master.zip -o master.zip
	@unzip -qq -o master.zip
	@cp -r iot-master/* ../
	@rm -rf iot-master master.zip
	@echo 'finish app install'
endif

setup:
	@pip install .

kill:
ifneq ("${ZOMBIES}","")
	@echo 'killing zombie process: ${ZOMBIES}'
	@kill -9 ${ZOMBIES}
else
	@echo 'nothing to kill'
endif

start: setup
	@echo 'starting app...'
	@python3 src/ds_sensors_mqtt/main.py

supervisor:
	@supervisord -c supervisord.conf

supervisor-program: install setup kill start

mosquitto:
	@docker run --rm -it -p 1883:1883 -p 9001:9001 \
		-v $(PWD)/mosquitto.conf:/mosquitto/config/mosquitto.conf \
		--name mosquitto eclipse-mosquitto
