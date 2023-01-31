package structs

import "encoding/json"

type Result interface {
	STR() string
	JSON() string
	SUCCESS() bool
	Get() *PocResult
}

type PocResult struct {
	Str            string
	Success        bool
	URL            string   `json:"url"`
	PocName        string   `json:"poc_name"`
	PocLink        []string `json:"poc_link"`
	PocAuthor      string   `json:"poc_author"`
	PocDescription string   `json:"poc_description"`
	Res            string   `json:"res"`
	Req            string   `json:"req"`
}

func (r *PocResult) JSON() string {
	if js, err := json.Marshal(r); err == nil {
		return string(js)
	}
	return ""
}

func (r *PocResult) STR() string {
	return r.Str
}

func (r *PocResult) SUCCESS() bool {
	return r.Success
}

func (r *PocResult) Get() *PocResult {
	return r
}
