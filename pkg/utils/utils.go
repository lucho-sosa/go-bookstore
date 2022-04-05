package utils

import (
	//import  encodings/json, io, and net/http
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}){
	//read the body of the request
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//parse the body into the v interface
		if err := json.Unmarshal([]byte(body), v); err != nil {
			return
		}
	}
}
