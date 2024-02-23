package homeassistant

import "fmt"

const (
	availabilityTopicFormat = "%s/bridge/state"
)

func (a *homeassistant) buildAvailability() {
	a.availabilityTopic = fmt.Sprintf(availabilityTopicFormat, a.bridge)
	a.availabilityPayload = "online"
}
