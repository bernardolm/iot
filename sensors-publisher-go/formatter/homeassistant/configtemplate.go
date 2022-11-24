package homeassistant

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

var (
	r = regexp.MustCompile(`[\s\p{Zs}]+|[\s\p{Zs}]+`)
)

type configPayloadTemplate struct {
	AvailabilityTopic string
	ConfigTopic       string
	DeviceClass       string
	ID                string
	Manufacturer      string
	Model             string
	Name              string
	StateTopic        string
	Unique            string
	UnitOfMeasurement string
}

func buildConfigPayload(t configPayloadTemplate) string {
	tmp1 := template.New("configPayloadTemplate").Delims("[[", "]]")
	payload := `
{
	'availability': [
		{
			'topic': '[[.AvailabilityTopic]]',
		}
	],
	'device': {
		'identifiers': [
			'[[.ID]]',
		],
		'manufacturer': '[[.Manufacturer]]',
		'model': '[[.Model]]',
		'name': '[[.Name]]',
		'sw_version': '[[.Model]] 0.0.1'
	},
	'device_class': '[[.DeviceClass]]',
	'json_attributes_topic': [[.StateTopic]],
	'name': '[[.Name]] ([[.ID]])',
	'state_class': 'measurement',
	'state_topic': [[.StateTopic]],
	'unique_id': [[.ID]],
	'unit_of_measurement': '[[.UnitOfMeasurement]]',
	'value_template': '{{ value_json.[[.DeviceClass]] }}',
}`
	payload = r.ReplaceAllString(payload, "")
	tmp1, err := tmp1.Parse(payload)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmp1.Execute(buf, t); err != nil {
		panic(err)
	}

	return buf.String()
}

func buildStatePayload(deviceClass string) string {
	return fmt.Sprintf(`{'%s': '%%v'}`, deviceClass)
}
