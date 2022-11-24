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

var (
	configPayload = `{
    "availability": [
      {
        "topic": "[[.AvailabilityTopic]]/state",
      },
      {
        "topic": "[[.StateTopic]]/availability",
      }
    ],
    "availability_mode": "all",
    "device": {
      "identifiers": [
        "[[.ID]]"
      ],
      "manufacturer": "[[.Manufacturer]]",
      "model": "[[.Model]]",
      "name": "[[.Name]]",
      "sw_version": "[[.Model]] 0.0.1"
    },
    "device_class": "[[.DeviceClass]]",
    "enabled_by_default": true,
    "json_attributes_topic": "[[.StateTopic]]",
    "name": "[[.Name]] ([[.ID]])",
    "state_class": "measurement",
    "state_topic": "[[.StateTopic]]",
    "unique_id": "[[.UniqueID]]",
    "unit_of_measurement": "[[.UnitOfMeasurement]]",
    "value_template": "{{ value_json.[[.DeviceClass]] }}"
	}`

	// spaceClear = regexp.MustCompile(`[\s\p{Zs}]+|[\s\p{Zs}]+`)
)

func (a *homeassistant) buildConfig(s sensor.Sensor) error {
	a.configTopic = fmt.Sprintf(configTopicFormat, name, s.ID(), s.DeviceClass())

	// payload := spaceClear.ReplaceAllString(configPayload, "")
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
		"Manufacturer":      s.Manufacturer(),
		"Model":             s.Model(),
		"Name":              fmt.Sprintf("%s %s sensor", s.Model(), s.DeviceClass()),
		"StateTopic":        a.stateTopic,
		"UnitOfMeasurement": s.UnitOfMeasurement(),
		"UniqueID":          fmt.Sprintf("%s_%s_%s", s.ID(), s.DeviceClass(), a.bridge),
	}

	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, data); err != nil {
		return err
	}

	a.configPayload = buf.String()

	return nil
}
