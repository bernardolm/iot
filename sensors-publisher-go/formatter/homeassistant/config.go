package homeassistant

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

const (
	configTopicFormat = "%s/sensor/%s/%s/config"
)

var configPayload = `
{
  "availability": [
    {
      "topic": "[[.AvailabilityTopic]]"
    },
    {
      "topic": "[[.StateTopic]]/availability"
    }
  ],
  "availability_mode": "all",
  "device": {
    "identifiers": [
      "[[.Identifier]]"
    ],
    "manufacturer": "[[.Manufacturer]]",
    "model": "[[.Model]]",
    "name": "[[.Name]]",
    "sw_version": "0.0.2"
  },
  "device_class": "[[.DeviceClass]]",
  "enabled_by_default": true,
  "json_attributes_topic": "[[.StateTopic]]",
  "name": "[[.Name]]",
  "state_class": "measurement",
  "state_topic": "[[.StateTopic]]",
  "unique_id": "[[.UniqueID]]",
  "unit_of_measurement": "[[.UnitOfMeasurement]]",
  "value_template": "{{ value_json.[[.DeviceClass]] }}"
}`

func (a *homeassistant) buildConfig(s sensor.Sensor) error {
	a.configTopic = fmt.Sprintf(configTopicFormat, name, s.ID(), s.DeviceClass())

	payload := configPayload

	configTempl := template.New("configPayloadTemplate").Delims("[[", "]]")
	templ, err := configTempl.Parse(payload)
	if err != nil {
		return err
	}

	data := map[string]string{
		"AvailabilityTopic": a.availabilityTopic,
		"DeviceClass":       s.DeviceClass(),
		"ID":                s.ID(),
		"Identifier":        fmt.Sprintf("%s_%s", a.bridge, s.ID()),
		"Manufacturer":      s.Manufacturer(),
		"Model":             s.Model(),
		"Name":              s.Name(),
		"StateTopic":        a.stateTopic,
		"UniqueID":          fmt.Sprintf("%s_%s", s.UniqueID(), a.bridge),
		"UnitOfMeasurement": s.UnitOfMeasurement(),
	}

	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, data); err != nil {
		return err
	}

	a.configPayload = buf.String()

	return nil
}
