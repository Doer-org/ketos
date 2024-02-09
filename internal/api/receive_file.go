package api

import (
	"io"
	"net/http"
	"os"
)

// TODO: 取り敢えず書いてみただけなので、動作確認しつつ適宜変更してください。
func ReceiveTarFromServer(id string) error {
	response, err := http.Get(BackendURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

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
