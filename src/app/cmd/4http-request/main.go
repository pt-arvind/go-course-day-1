package main

import (
	"app/internal/constant"
	"app/internal/gateway"
	"app/internal/model"
	"encoding/json"
	"fmt"
	"log"
)

/*
	Make an HTTP request through the code to http://www.omdbapi.com/?apikey=<YOUR API KEY>&t=matrix
	and parse out the JSON into your struct from step 2
	print the struct to the screeen
*/

func main() {
	respBytes, err := gateway.MakeOMDBRequest("the matrix", constant.OMDBURL, constant.APIKey)
	fmt.Printf("%+v\n", string(respBytes))
	if err != nil {
		log.Printf("error performing request: %+v", err)
		return
	}

	response := model.Movie{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		log.Printf("Error: %+v", err)
		return
	}
	log.Printf("Response: %+v", response)
}
