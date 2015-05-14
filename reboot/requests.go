package reboot

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"io"
)

const PENDING_URL = "/api/pending/"

type Request struct {
	Url               string
	Family_name       string
	Given_name        string
	Requester_type    string
	Faculty_and_dept  string
	Organization      string
	Preset            string
	presetObject	*Preset
	Os                string
	osObject	*OS
	Machine_use       string
	Need_display      bool
	Need_keyboard     bool
	Need_mouse        bool
	Need_ethernet     bool
	Extra_information string
	Amount            int
	Filled            bool
	Filled_at         *time.Time
	Requested_at      *time.Time
}

func (r *Request) Fetch() *Request {
	page, err := http.Get(r.Url+JSON_FORMAT)
	if err != nil {
		log.Fatal(err)
	}
	defer page.Body.Close()
	dec := json.NewDecoder(page.Body)
	if err = dec.Decode(r); err == io.EOF {
		log.Fatal(err)
	} else if err != nil {
		log.Println("Error fetching Request at ", r.Url)
		log.Fatal(err)
	}
	return r
}
