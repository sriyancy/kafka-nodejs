package controllers

import ("net/http"
"io"
"encoding/json")

func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
func encodeResponseASJSON(data interface{},w io.Writer){
	enc:=json.NewEncoder(w)
	enc.Encode(data)
}
