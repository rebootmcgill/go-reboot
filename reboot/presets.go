package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type Preset struct {
	Url string
	cpu CPU
	Ram int
	Hdd int
}
