package main

import (
	"encoding/json"
	"io"
	"net/http"
	us "user"
)

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func selectAll(w http.ResponseWriter, r *http.Request) {
	usr := us.SelectAll()
	data, err := json.Marshal(usr)
	cekError(err)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tampilUserAll", selectAll)

	http.ListenAndServe(":5050", mux)
}
