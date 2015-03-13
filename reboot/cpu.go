package reboot

import (
	_ "encoding/json"
	_ "net/http"
)

type CPU struct {
	url   string
	name  string
	cores int
	x64   bool
	clock float64
}
