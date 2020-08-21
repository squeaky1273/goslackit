package slack

import (
	// "encoding/json"
	// "log"
	"net/http"
	"io/ioutil"
)

// type jokes struct {
// 	Setup      string  `json:"setup"`
// 	Punchline  string  `json:"punchline"`
// }

func DadJokes() (string, error) {
	url := "https://dad-jokes.p.rapidapi.com/random/jokes"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "dad-jokes.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "JOKE_API")
	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// joke := jokes{}

	// jsonErr := json.Unmarshal(body, &joke)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	// return joke, err

	return string(body), err

}