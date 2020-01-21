package cmd

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type DBOpenerFactory struct {}

type DBOpener interface {
	open(c DBCredentials) (err error)
}

func (f *DBOpenerFactory) make(app string, ui cli.Ui) (o DBOpener, err error) {
	switch app {
	case "tableplus":
		return &DBOpenerTableplus{}, nil
	case "sequel-pro":
		return &DBOpenerSequelPro{ui: ui}, nil
	}

	return nil, fmt.Errorf("Error: %s is not supported", app)
}
