package runner

import (
	"os/exec"
	"strings"

	"github.com/bernardolm/iot/supervisor-go/config"
	"github.com/bernardolm/iot/supervisor-go/installer"
	log "github.com/sirupsen/logrus"
)

func Run(p []config.Program) error {
	for k := range p {
		sCmd := strings.Fields(p[k].Command)
		var cmd *exec.Cmd

		if p[k].NeedInstall() {
			log.Debugf("installing program '%s' with url '%s'\n", p[k].Name, p[k].URL)

			path, err := installer.Install(p[k])
			if err != nil {
				return err
			}
			cmd = exec.Command(path, sCmd[1:]...)
		} else {
			cmd = exec.Command(sCmd[0], sCmd[1:]...)
		}

		// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		log.Debugf("starting program '%s' with command '%s'\n", p[k].Name, p[k].Command)

		err := cmd.Start()
		if err != nil {
			return err
		}

		p[k].PID = cmd.Process.Pid

		go func(n string) {
			err = cmd.Wait()
			log.Debugf("[%s] %v\n", n, err)
		}(p[k].Name)
	}

	return nil
}
