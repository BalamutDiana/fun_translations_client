package funtranslations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Languages struct {
	Pirate string
}

func LanguagesList() *Languages {
	return &Languages{
		Pirate: "pirate",
	}
}

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {

	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetTranslation(language, text string) (ResponseData, error) {
	values := map[string]string{"text": text}
	json_data, err := json.Marshal(values)

	if err != nil {
		return ContentsData{}, err
	}

	url := fmt.Sprintf("https://api.funtranslations.com/translate/%s.json", language)
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return ContentsData{}, err
	}
	defer resp.Body.Close()

	available, err := checkAvailability(resp)
	if err != nil {
		return ContentsData{}, err
	}

	if available {
		var r assetsResponse

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return r.Contents, err
		}

		err = json.Unmarshal(body, &r)
		if err != nil {
			return r.Contents, err
		}

		return r.Contents, nil

	} else {

		var r assetsError

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return r.Contents, err
		}

		err = json.Unmarshal(body, &r)
		if err != nil {
			return r.Contents, err
		}

		return r.Contents, nil
	}

}

func checkAvailability(resp *http.Response) (bool, error) {

	switch resp.StatusCode {
	case 200:
		return true, nil
	case 429:
		return false, nil
	default:
		return false, errors.New("something went wrong, check the request or try again later")
	}

}
