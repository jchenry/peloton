package peloton

import (
	"bytes"
	"fmt"

	"github.com/jchenry/peloton/model"
)

const bookmarkPayload = "{\"ride_id\":\"%s\"}"

//CreateBookmark creates a bookmark for a specified ride for the user of the client
func CreateBookmark(c Client, r model.Ride) bool {
	resp, _ := c.HTTPClient.Post(fmt.Sprintf(
		"%s/api/favorites/create", apiBase),
		contentTypeJSON,
		bytes.NewReader([]byte(fmt.Sprintf(bookmarkPayload, r.ID))))
	return resp.StatusCode == 200
}

//RemoveBookmark removes a bookmark for a specified ride for the user of the client
func RemoveBookmark(c Client, r model.Ride) bool {
	resp, _ := c.HTTPClient.Post(fmt.Sprintf(
		"%s/api/favorites/delete", apiBase),
		contentTypeJSON,
		bytes.NewReader([]byte(fmt.Sprintf(bookmarkPayload, r.ID))))
	return resp.StatusCode == 200
}
