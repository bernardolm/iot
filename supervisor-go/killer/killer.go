package killer

import (
	"fmt"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/bernardolm/iot/supervisor-go/config"
	log "github.com/sirupsen/logrus"
)

func kill(pid int) error {
	err := syscall.Kill(-pid, syscall.SIGKILL)
	if err != nil {
		return err
	}

	log.Debugf("%v killed\n", pid)
	return nil
}

func Init(p []config.Program) {
	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if len(q) > 0 {
			log.Debug(q)
			p := q.Get("pid")

			i, err := strconv.Atoi(p)
			if err != nil {
				log.Error(err)
				return
			}
			if err := kill(i); err != nil {
				log.Error(err)
				return
			}
		}

		for k := range p {
			if p[k].PID == 0 {
				log.Debugf("program '%s' already killed\n", p[k].Name)
				continue
			}

			if err := kill(p[k].PID); err != nil {
				log.Error(err)
			}
			p[k].PID = 0
		}

		if _, err := w.Write([]byte(fmt.Sprintf("killed at %s", time.Now()))); err != nil {
			log.Error(err)
		}
	})
}
