package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type User struct {
	Url      string
	Username string
	Is_staff bool
}
