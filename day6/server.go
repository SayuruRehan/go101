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
			res.Text = fmt.Sprintf("Hello, %s", who)
		}

		jsonRes, _ := json.Marshal(res)
		io.WriteString(w, string(jsonRes))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}


/*
EXPLANATION:

1. Package Declaration: The code starts with the declaration of the main package, which is the entry point for a Go application.
2. Imports: The necessary packages are imported:
   - "fmt" for formatted I/O operations.
   - "net/http" for HTTP server and client implementations.
   - "encoding/json" for JSON encoding and decoding.
   - "io" for basic I/O operations.
   - "log" for logging errors.
3. Main Function: The main function is defined, which sets up the HTTP server.
4. HTTP Handler: An HTTP handler function is registered for the "/hello" endpoint using `http.HandleFunc`. This function will be called whenever a request is made to this endpoint.
5. Greeting Struct: Inside the handler function, a struct named `Greeting` is defined to represent the JSON response structure. It has two fields: `Text` for the greeting message and `Error` for any error messages (the `omitempty` tag ensures
   that the `Error` field is omitted from the JSON if it is empty).
6. Query Parameter Extraction: The handler extracts the "who" query parameter from the request URL using `r.URL.Query().Get("who")`.
7. Response Preparation: An instance of the `Greeting` struct is created. If the "who" parameter is empty, an error message is set in the `Error` field. Otherwise, a greeting message is formatted and assigned to the `Text` field.
8. JSON Encoding: The response struct is marshaled into JSON format using `json.Marshal`.
9. Writing Response: The JSON response is written to the HTTP response writer using `io.WriteString`.
10. Starting the Server: The HTTP server is started on port 8080 using `http.ListenAndServe`. If there is an error starting the server, it is logged using `log.Println`.
This code sets up a simple HTTP server that responds with a JSON greeting message based on the "who" query parameter provided in the request.
*/
