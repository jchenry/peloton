package peloton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiBase = "https://api.onepeloton.com"

// Client stores all state for communicating with the Peloton API
type Client struct {
	HTTPClient http.Client
	user       user
}

/*
	j, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c := peloton.Client {
		HTTPClient: http.Client{
			Jar: j
		}
	}

	Authenticate(c)
*/

type user struct {
	ID string `json:"user_id"`
}

//Authenticate authenticates a client with a username and password
func Authenticate(client *Client, username, password string) error {
	payload := map[string]string{
		"username_or_email": username,
		"password":          password,
	}

	j, _ := json.Marshal(payload)

	resp, _ := client.HTTPClient.Post(
		fmt.Sprintf("%s/auth/login", apiBase),
		contentTypeJSON,
		bytes.NewReader(j))
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		u := user{}
		json.Unmarshal(body, &u)
		client.user = u
	}

	return err
}
