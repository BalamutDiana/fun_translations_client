package main

import (
	"fmt"
	"golang-ninja/httpclient/funtranslations"
	"log"
	"time"
)

func main() {
	funtranslationsClient, err := funtranslations.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	trans, err := funtranslationsClient.GetTranslation()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(trans.Info())
}
