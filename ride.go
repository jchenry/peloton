package peloton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/jchenry/peloton/model"
)

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

//SearchFilter represents the data needed to filter Peloton's ride catalog
type SearchFilter struct {
	Category        RideCategory //browse_category
	Duration        int          //duration (sec)
	Difficulty      string       //difficulty_level
	Booknmarked     bool         //is_favorite_ride (Bookmarked)
	ClosedCaptioned bool         //has_closed_captions
	ClassType       string       //class_type_id
	MusicType       string       //super_genre_id
	Instructor      string       //instructor_id
	Page            int          //page
	Limit           int          //limit
	SortBy          string       //sort_by (desc)
}

func (f *SearchFilter) toURLQuery() string {
	v := url.Values{}
	v.Set("browse_category", f.Category.String())
	if f.Duration > 0 {
		v.Set("duration", strconv.Itoa(f.Duration*60))
	}
	if !strings.EqualFold(f.Difficulty, "") {
		v.Set("difficulty_level", f.Difficulty)
	}

	if f.Booknmarked {
		v.Set("is_favorite_ride", strconv.FormatBool(f.Booknmarked))
	}
	if f.ClosedCaptioned {
		v.Set("has_closed_captions", strconv.FormatBool(f.ClosedCaptioned))
	}

	v.Add("content_format", "audio")
	v.Add("content_format", "video")

	if !strings.EqualFold(f.ClassType, "") {
		v.Set("class_type_id", f.ClassType)
	}
	if !strings.EqualFold(f.MusicType, "") {
		v.Set("super_genre_id", f.MusicType)
	}
	if !strings.EqualFold(f.Instructor, "") {
		v.Set("instructor_id", f.Instructor)
	}

	if f.Page > 0 {
		v.Set("page", strconv.Itoa(f.Page))
	}

	if f.Limit > 0 {
		v.Set("limit", strconv.Itoa(f.Limit))
	}

	return v.Encode()

}

//GetRides return a list of Rides based on the SearchFilter
func GetRides(c Client, f SearchFilter) ([]model.Ride, error) {
	a, err := getArch(c, f)
	return a.Rides, err
}

//GetRideTypes return a list of RideTypes based on the SearchFilter
func GetRideTypes(c Client, f SearchFilter) ([]model.RideType, error) {
	a, err := getArch(c, f)
	return a.RideTypes, err
}

//GetInstructors return a list of Instructors based on the SearchFilter
func GetInstructors(c Client, f SearchFilter) ([]model.Instructor, error) {
	a, err := getArch(c, f)
	return a.Instructors, err
}

func getArch(c Client, f SearchFilter) (*model.Archive, error) {
	arch := &model.Archive{}

	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/api/v2/ride/archived%s", apiBase, f.toURLQuery()))
	if err != nil {
		return arch, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &arch)
	return arch, err
}
