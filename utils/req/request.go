package req

import (
	"adventofcode2021/secrets"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MakeRequest make request to advent of code
func MakeRequest(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", secrets.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
