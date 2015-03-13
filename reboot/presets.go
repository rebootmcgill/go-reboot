package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type Preset struct {
	url string
	cpu CPU
	ram int
	hdd int
}
