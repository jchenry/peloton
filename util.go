package peloton

import (
	"github.com/jchenry/peloton/model"
)

var contentTypeJSON = "application/json"

//PrettyInstructor returns the "pretty" string for a instructorID from a list of given Instructors
func PrettyInstructor(instructors []model.Instructor, instructorID string) string {
	for _, instructor := range instructors {
		if instructor.ID == instructorID {
			return instructor.Name
		}
	}
	return "Unknown " + instructorID
}

//PrettyRideType returns the "pretty" version for a rideTypeID from a list of given RideTypes
func PrettyRideType(rideTypes []model.RideType, rideTypeID string) string {
	for _, rideType := range rideTypes {
		if rideType.ID == rideTypeID {
			return rideType.Name
		}
	}
	return "Unknown"
}
