package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO: 取り敢えず書いてみただけなので、動作確認しつつ適宜変更してください。
func ReceiveTarGzFromServer(id string) error {
	fullURL := fmt.Sprintf("%s/%s", BackendURL, id)
	response, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	infoURL := fmt.Sprintf("%s/info/%s", BackendURL, id)
	infoResponse, err := http.Get(infoURL)
	if err != nil {
		return err
	}
	defer infoResponse.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, infoResponse.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody.String()))

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
