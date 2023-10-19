package influxdb

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Start(ctx context.Context) {
	log.Info("influxdb: starting client")
	_ = getAPI(ctx)
}

func Send(ctx context.Context, _ string, payload interface{}) {
	log.Debug("influxdb: publishing")

	line := payload.([]byte)
	getAPI(ctx).WriteRecord(string(line))
	getAPI(ctx).Flush()

	log.
		WithField("payload", fmt.Sprintf("%s", payload)).
		Info("influxdb: sent (writed)")
}

func Finish(ctx context.Context) {
	log.Debug("influxdb: disconnecting")

	api = nil
	getClient(ctx).Close()
	client = nil

	log.Info("influxdb: disconnected")
}
