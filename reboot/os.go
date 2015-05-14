package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type OS struct {
	Url          string
	Name         string
	Version      string
	Experimental bool
}
