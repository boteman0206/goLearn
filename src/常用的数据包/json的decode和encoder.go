package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	js := `{""name":"jack","html": "<p>p标签</p>"}`
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(js)

	fmt.Println("jac : ", err, buf)
	fmt.Println(buf.String())

}
