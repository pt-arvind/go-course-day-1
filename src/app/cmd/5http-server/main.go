package main

import (
	"app/internal/constant"
	"app/internal/gateway"
	"app/internal/model"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

/*
	Take the main function in the last step and rename/move it to the internal folder package
	Start a webserver on port 9000
	Set up a route: GET /movie/search which calls the function described above and returns the result
*/
func main() {
	http.HandleFunc("/movie/search", func(w http.ResponseWriter, r *http.Request) {
		movieTitle := r.URL.Query().Get("t")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if movieTitle == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "{\"error\":\"Bad Request, title is a required query parameter\"}")
			return
		}

		respBytes, err := gateway.MakeOMDBRequest(movieTitle, constant.OMDBURL, constant.APIKey)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "{\"error\":\"Internal Server Error requesting OMDB Server: %+v\"}", err)
			return
		}

		response := model.Movie{}
		err = json.Unmarshal(respBytes, &response)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "{\"error\":\"Internal Server Error unmarshaling OMDB response: %+v\"}", err)
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "{\"error\":\"Internal Server Error marshaling server response: %+v\"}", err)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404)
		fmt.Fprintf(w, "Hello, %q does not exist", html.EscapeString(r.URL.Path))
	})

	port := 9000
	log.Printf("Starting server on: localhost:%v...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
