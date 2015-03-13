package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type OS struct {
	url          string
	name         string
	version      string
	experimental bool
}
