package tmdb

import (
	"encoding/json"
	"io"
	"net/http"
)

func SearchTMDB(title, year, apiKey, lang string) (map[string]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/search/movie", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("query", title)
	q.Add("year", year)
	q.Add("language", lang)
	q.Add("api_key", apiKey)
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
	if results, ok := result["results"].([]interface{}); ok && len(results) > 0 {
		firstResult := results[0].(map[string]interface{})
		foundTitle := firstResult["title"].(string)
		foundYear := firstResult["release_date"].(string)[:4]
		ret = map[string]string{
			"title": foundTitle,
			"year":  foundYear,
		}
		return ret, nil
	}
	return ret, nil
}
