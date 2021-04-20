package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	//register handler
	//er can use for ntwok req
	//return static txt wen i req from app
	//"/"route path every req
	//func are 1st clas citizens
	//ananymous func no need to preinitialize
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)

		//w.Write([]byte("Hello " + name))
	})
	err := http.ListenAndServe(":5006", nil)
	if err != nil {
		log.Fatal(err)
	}

}
