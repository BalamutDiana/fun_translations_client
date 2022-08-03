package funtranslations

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

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

func (c Client) GetTranslation() (ContentsData, error) {
	values := map[string]string{"text": "Hello little bro"}
	json_data, err := json.Marshal(values)

	if err != nil {
		return ContentsData{}, err
	}

	resp, err := c.client.Post("https://api.funtranslations.com/translate/pirate.json", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return ContentsData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ContentsData{}, err
	}

	//fmt.Println(string(body))
	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return ContentsData{}, err
	}

	return r.Contents, nil
	//fmt.Println(r.Contents.Info())
}
