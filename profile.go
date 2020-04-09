package peloton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Profile represents the vital stats of a user of peloton
type Profile struct {
	Weight float32 `json:"weight"`
}

//GetProfile returns the profile for the user of the client
func GetProfile(c Client) (Profile, error) {
	p := Profile{}

	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/api/user/%s", apiBase, c.user.ID))
	if err != nil {
		return p, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(body, &p)
	return p, err
}
