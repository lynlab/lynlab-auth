package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type googleUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func getGoogleUser(token string) (u *googleUser, e error) {
	req, _ := http.NewRequest("GET", "https://oauth2.googleapis.com/tokeninfo", nil)
	q := req.URL.Query()
	q.Add("id_token", token)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		e = err
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &u)
	return
}
