package reboot

import (
	_ "encoding/json"
	_ "net/http"
	"time"
)

type Machine struct {
	url         string
	request     Request
	fulfiller   User
	cpu         CPU
	ram         int
	hdd         int
	pickedup_at time.Time
	notes       string
	picked_up   bool
}
