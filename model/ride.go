package model

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
