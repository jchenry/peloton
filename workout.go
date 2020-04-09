package peloton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jchenry/peloton/model"
)

//GetWorkouts returns a workout history for the authenticated user
func GetWorkouts(c Client, limit, page int) (model.WorkoutHistory, error) {
	resp, _ := c.HTTPClient.Get(fmt.Sprintf("%s/api/user/%s/workouts?joins=ride&limit=%d&page=%d", apiBase, c.user.ID, limit, page))

	wh := model.WorkoutHistory{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wh, err
	}

	err = json.Unmarshal(body, &wh)
	return wh, err
}
