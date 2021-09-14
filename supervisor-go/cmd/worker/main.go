package main

import (
	"fmt"
	"net/http"

	"github.com/bernardolm/iot/supervisor-go/config"
	"github.com/bernardolm/iot/supervisor-go/github"
	"github.com/bernardolm/iot/supervisor-go/killer"
	"github.com/bernardolm/iot/supervisor-go/runner"
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.SetLevel(log.DebugLevel)

	cfg, err := config.Load()
	if err != nil {
		log.Panic(err)
	}
	log.Debug(pp.Sprint(cfg))

	if err := runner.Run(cfg.Programs); err != nil {
		log.Panic(err)
	}
	log.Debug(pp.Sprint(cfg))

	if err := github.AddRoutes(cfg.Github.Secret); err != nil {
		log.Panic(err)
	}

	killer.Init(cfg.Programs)

	fmt.Printf("serving in port %s\n", viper.GetString("http.port"))
	log.Panic(http.ListenAndServe(":"+viper.GetString("http.port"), nil))
}
