package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 5,
	}

	values := map[string]string{"text": "Hello little bro"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Post("https://api.funtranslations.com/translate/pirate.json", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
