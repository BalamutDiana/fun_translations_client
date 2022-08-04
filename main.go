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

	lang := funtranslations.LanguagesList()
	text := "Can you tell me how to get to the library?"

	trans, err := funtranslationsClient.GetTranslation(lang.Pirate, text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trans.Info())

}
