package omdb

import (
	"encoding/json"
	"io"
	"net/http"
)

func SearchOMDB(title, year, apiKey string) (map[string]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.omdbapi.com/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("t", title)
	q.Add("y", year)
	q.Add("apikey", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	ret := map[string]string{}
	if result["Response"] == "True" {
		ret = map[string]string{
			"title": result["Title"].(string),
			"year":  result["Year"].(string),
		}
		return ret, nil
	}

	return ret, nil
}
