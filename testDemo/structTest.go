package easyTest

import "time"

//easyjson:json
type School struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

//easyjson:json
type Student struct {
	Id       int       `json:"id"`
	Name     string    `json:"s_name"`
	School   School    `json:"s_chool"`
	Birthday time.Time `json:"birthday"`
}
