package model

// http://www.omdbapi.com/?t=matrix&y=1999

/*
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
}
*/

// Movie represents a movie in the omdb
type Movie struct {
	Title          string        `json:"Title"`
	Year           int           `json:"Year,string"`
	Rated          string        `json:"Rated"`
	Released       string        `json:"Released"`
	Runtime        string        `json:"Runtime"`
	Genre          string        `json:"Genre"`
	Director       string        `json:"Director"`
	Writer         string        `json:"Writer"`
	Actors         string        `json:"Actors"`
	Plot           string        `json:"Plot"`
	Language       string        `json:"Language"`
	Country        string        `json:"Country"`
	Awards         string        `json:"Awards"`
	Poster         string        `json:"Poster"`
	Ratings        []MovieRating `json:"Ratings"`
	Metascore      string        `json:"Metascore"`
	IMDBRating     string        `json:"imdbRating"`
	IMDBVotes      string        `json:"imdbVotes"`
	IMDBID         string        `json:"imdbID"`
	Type           string        `json:"Type"`
	DVDReleaseDate string        `json:"DVDReleaseDate"`
	BoxOffice      string        `json:"BoxOffice"`
	Production     string        `json:"Production"`
	Website        string        `json:"Website"`
	Response       string        `json:"Response"`
}

// MovieRating is the movie rating
type MovieRating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
