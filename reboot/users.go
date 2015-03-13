package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type user struct {
	url      string
	username string
	is_staff bool
}
