package main

import (
	"app/internal/model"
	"encoding/json"
	"log"
)

/*
	Read the JSON string below and turn it into the sunset time struct we just made in step 2
	print the struct to the screeen
	Copy the struct to an 'internal' folder outside the 'cmd' folder to preserve golang conventions
*/
var jsonString = `
{
	"Title": "The Matrix",
	"Year": "1999",
	"Rated": "R",
	"Released": "31 Mar 1999",
	"Runtime": "136 min",
	"Genre": "Action, Sci-Fi",
	"Director": "Lana Wachowski, Lilly Wachowski",
	"Writer": "Lilly Wachowski, Lana Wachowski",
	"Actors": "Keanu Reeves, Laurence Fishburne, Carrie-Anne Moss, Hugo Weaving",
	"Plot": "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
	"Language": "English",
	"Country": "USA",
	"Awards": "Won 4 Oscars. Another 34 wins & 48 nominations.",
	"Poster": "https://m.media-amazon.com/images/M/MV5BNzQzOTk3OTAtNDQ0Zi00ZTVkLWI0MTEtMDllZjNkYzNjNTc4L2ltYWdlXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_SX300.jpg",
	"Ratings": [{
		"Source": "Internet Movie Database",
		"Value": "8.7/10"
	}, {
		"Source": "Rotten Tomatoes",
		"Value": "88%"
	}, {
		"Source": "Metacritic",
		"Value": "73/100"
	}],
	"Metascore": "73",
	"imdbRating": "8.7",
	"imdbVotes": "1,549,931",
	"imdbID": "tt0133093",
	"Type": "movie",
	"DVD": "21 Sep 1999",
	"BoxOffice": "N/A",
	"Production": "Warner Bros. Pictures",
	"Website": "N/A",
	"Response": "True"
}`

func main() {
	response := model.Movie{}
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		log.Printf("Error: %+v", err)
		return
	}
	log.Printf("Response: %+v", response)
}
