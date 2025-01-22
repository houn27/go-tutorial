package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

func StructAss() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	for _, loc := range locations {
		loc_json, err := json.Marshal(loc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(loc_json))
	}

	// loc_json_beau, err := json.MarshalIndent(locations)
	// if err != nil {
	// 	fmt.Println(loc_json)
	// }

}

// ------------------------- Method Ass --------------------

type location struct {
	rover, landing string
	lat, long      float64
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(rover, landing string, lat, long coordinate) location {
	return location{rover, landing, lat.decimal(), long.decimal()}
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// stringfy location

func (l location) description() string {
	return fmt.Sprintf("%+v", l)
}

// ------------------------- Composition Assignment --------------------

type new_location struct {
	landing   string
	lat, long float64
}

type rover struct {
	name string
	gps
}

type gps struct {
	cur_loc, dest_loc new_location
	world
}

type world struct {
	radius float64
}

// Method to calculate the distance between two locations using the spherical law of cosines
func (w world) distance(p1, p2 new_location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// Helper function to convert degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) massege() string {
	return fmt.Sprintf("%f km left", g.distance(g.cur_loc, g.dest_loc))
}

// ------------------------- Interface Assignment --------------------

func (c coordinate) String() string {
	sign := 1.0
	switch c.h {
	case 'S', 'W':
		sign = -1
	}

	return fmt.Sprintf("%v%v∘%v'%v''%v", sign, c.d, c.m, c.s, string(c.h))
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	type coo_json_format struct {
		Demical float64 `json:"decimal"`
		Dms     string  `json:"dms"`
		D       float64 `json:"degrees"`
		M       float64 `json:"minutes"`
		S       float64 `json:"seconds"`
		H       rune    `json:"hemisphere"`
	}

	coo_json := coo_json_format{c.decimal(), fmt.Sprintf("%v", c), c.d, c.m, c.s, c.h}

	return json.Marshal(coo_json)
}

func main() {

	// run method ass
	data := []location{
		newLocation("Spirit", "Columbia Memorial Station", coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'}),
		newLocation("Opportunity", "Challenger Memorial Station", coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'}),
		newLocation("Curiosity", "Bradbury Landing", coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'}),
		newLocation("InSight", "Elysium Planitia", coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'}),
	}

	for _, i := range data {
		fmt.Println(i.description())
	}

	// run Composition Assignment
	new_gps := gps{new_location{"Bradbury Landing", -4.58995, 127.4417}, new_location{"Elysium Planitia", 2.5, 135.9}, world{3389.5}}
	rover := rover{"Curoisity", new_gps}
	fmt.Println(rover.massege())

	// run Interface Assignment
	bytes, err := json.Marshal(coordinate{4, 35, 22.2, 'S'})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
