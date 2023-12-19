package installer

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/bernardolm/iot/supervisor-go/config"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
)

const (
	permission fs.FileMode = 0744
)

func Install(p config.Program) (string, error) {
	pPath := "./bin"

	if _, err := os.Stat(pPath); os.IsNotExist(err) {
		err := os.Mkdir(pPath, permission)
		if err != nil {
			fmt.Printf("%#v", err)
			return "", err
		}
	}

	pPath = path.Join(pPath, slug.Make(p.Command))

	log.Debugf("installing in %s\n", pPath)

	out, err := os.Create(pPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(p.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = out.Truncate(0)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(pPath, body, permission)
	if err != nil {
		return "", err
	}

	return pPath, nil
}
