package peloton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jchenry/peloton/model"
)

//GetRideTypes return a list of RideTypes based on the SearchFilter
func GetRideTypes(c Client, f SearchFilter) ([]model.RideType, error) {
	a, err := getArch(c, f)
	return a.RideTypes, err
}

//GetInstructors return a list of Instructors based on the SearchFilter
func GetInstructors(c Client) ([]model.Instructor, error) {
	m, err := metadataMappings(c)
	return m.Instructors, err
}

func metadataMappings(c Client) (*model.MetadataMappings, error) {
	meta := &model.MetadataMappings{}

	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/api/ride/metadata_mappings", apiBase))
	if err != nil {
		return meta, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &meta)
	return meta, err
}
