package funtranslations

import "fmt"

type assetsResponse struct {
	Success  successData  `json:"success"`
	Contents ContentsData `json:"contents"`
}

type ContentsData struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

func (c ContentsData) Info() string {
	return fmt.Sprintf("Translation from english to %s\n Source text: %s\n Translation: %s",
		c.Translation, c.Text, c.Translated)
}

type successData struct {
	Total int `json:"total"`
}
