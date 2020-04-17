package peloton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jchenry/peloton/model"
)

//GetProfile returns the profile for the user of the client
func GetProfile(c Client) (model.Profile, error) {
	p := model.Profile{}

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
