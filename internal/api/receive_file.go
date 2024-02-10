package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetServerInfo(id string) error {
	fullURL := fmt.Sprintf("%s/info/%s", BackendURL, id)
	response, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, response.Body)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%s/info/%s", FrontURL, id))
	fmt.Println(string(responseBody.String()))

	return nil
}

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

	err = GetServerInfo(id)
	if err != nil {
		return err
	}

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
