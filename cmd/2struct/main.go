package main

import (
	"fmt"
	"time"
)

/*
Visit https://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400 and model a struct off of this
Use fmt.Printf with the %+v option to print it out to screen
*/

// SunInfo defines sunrise-sunset.org API response
type SunInfo struct {
	Sunrise                   time.Time
	Sunset                    time.Time
	SolarNoon                 time.Time
	DayLength                 time.Duration
	CivilTwilightBegin        time.Time
	CivilTwilightEnd          time.Time
	NauticalTwilightBegin     time.Time
	NauticalTwilightEnd       time.Time
	AstronomicalTwilightBegin time.Time
	AstronomicalTwilightEnd   time.Time
}

func main() {
	fmt.Printf("%+v\n", SunInfo{})
}
