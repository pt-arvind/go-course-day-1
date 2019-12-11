package main

import (
	"app/internal/model"
	"fmt"
)

/*
Visit http://www.omdbapi.com/?i=tt3896198&apikey=a9f87da0&t=matrix&y=1999 and model a struct off of this
Use fmt.Printf with the %+v option to print it out to screen
*/

func main() {
	fmt.Printf("%+v\n", model.Movie{})
}
