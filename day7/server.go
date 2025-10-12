package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		type Greeting struct {
			Text  string `json:"text"`
			Error string `json:"error,omitempty"`
		}

		who := r.URL.Query().Get("who")
		res := Greeting{}
		if who == "" {
			res.Error = "Greet who?"
		} else {
			res.Text = fmt.Sprintf("Hello %s", who)
		}

		jsonRes, _ := json.Marshal(res)
		io.WriteString(w, string(jsonRes))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}