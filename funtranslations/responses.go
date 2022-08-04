package funtranslations

import "fmt"

type ResponseData interface {
	Info() string
}

type assetsResponse struct {
	Contents ContentsData `json:"contents"`
	Success  SuccessData  `json:"success"`
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

func (c assetsResponse) JsonStructure() assetsResponse {
	return c
}

type SuccessData struct {
	Total int `json:"total"`
}

type assetsError struct {
	Contents ErrorData `json:"error"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorData) Info() string {
	return e.Message
}
