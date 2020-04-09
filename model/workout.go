package model

type WorkoutHistory struct {
	Workouts []Workout `json:"data"`
}

type Workout struct {
	Status    string `json:"status"`
	StartTime int64  `json:"start_time"`
	TZ        string `json:"timezone"`
	Ride      Ride   `json:"ride"`
}
