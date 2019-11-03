package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	grant_type = "client_credentials"
	authURL    = "https://api.line.me/v2/oauth/accessToken"
)

type LINEToken struct {
	AccessToken string `json:"access_token"`
}

func GetAccessToken() (*LINEToken, error) {
	values := url.Values{}
	values.Set("grant_type", grant_type)
	values.Add("client_id", channelID)
	values.Add("client_secret", channelSecret)

	req, err := http.NewRequest(
		"POST",
		authURL,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, fmt.Errorf("Invalid Status: %s", string(body))
	}

	lt := &LINEToken{}
	json.Unmarshal(body, lt)

	return lt, nil
}
