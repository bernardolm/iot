package influxdb

import (
	"context"

	influxapi "github.com/influxdata/influxdb-client-go/api"
	log "github.com/sirupsen/logrus"
)

var api influxapi.WriteAPI

func getAPI(ctx context.Context) influxapi.WriteAPI {
	if api == nil {
		api = getClient(ctx).WriteAPI("dummy-org", database)

		errorsCh := api.Errors()
		go func() {
			for err := range errorsCh {
				log.Fatalf("influxdb: api write error - %s\n", err.Error())
			}
		}()
	}

	return api
}
