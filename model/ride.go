package model

import "strings"

// jq -r '.data[] | [ .id,.title,.description,.instructor_id,.duration,.difficulty_rating_avg ] | @csv' ~/Downloads/peloton-rides.json > ~/Downloads/peloton-rides.csv

type Archive struct {
	Rides       []Ride       `json:"data"`
	Instructors []Instructor `json:"instructors"`
	RideTypes   []RideType   `json:"ride_types"`
}

type Instructor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RideType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Ride struct {
	ID            string  `json:"id"`
	RideType      string  `json:"ride_type_id"`
	RecordingDate int64   `json:"original_air_time"`
	Title         string  `json:"title"`
	Instructor    string  `json:"instructor_id"`
	Description   string  `json:"description"`
	Duration      int     `json:"duration"`
	Difficulty    float64 `json:"difficulty_rating_avg"`
}

//RideCategory represents the category of ride to take in peloton
type RideCategory int

func (r RideCategory) String() string {
	return strings.ToLower([]string{
		"Outdoor",
		"Running",
		"Cycling",
		"Strength",
		"Yoga",
		"Meditation",
		"Stretching",
		"Bootcamp",
		"Walking",
		"Cardio",
	}[r])
}

//GetRideCategory returs the RideCategory based on the string representation
func GetRideCategory(s string) RideCategory {
	switch s {
	case "Outdoor":
		return Outdoor
	case "Running":
		return Running
	case "Cycling":
		return Cycling
	case "Strength":
		return Strength
	case "Yoga":
		return Yoga
	case "Meditation":
		return Meditation
	case "Stretching":
		return Stretching
	case "Bootcamp":
		return Bootcamp
	case "Walking":
		return Walking
	case "Cardio":
		return Cardio
	default:
		return Cycling
	}
}

const (
	Outdoor RideCategory = iota
	Running
	Cycling
	Strength
	Yoga
	Meditation
	Stretching
	Bootcamp
	Walking
	Cardio
)
