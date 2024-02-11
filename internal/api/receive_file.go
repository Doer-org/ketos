package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
    Port string `json:"port"`
}

func GetServerInfo(id string) (string,error) {
	fullURL := fmt.Sprintf("%s/info/%s", BackendURL, id)
	response, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, response.Body)
	if err != nil {
		return "", err
	}

	var resp Response
	err = json.Unmarshal(responseBody.Bytes(), &resp)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s/info/%s\n", FrontURL, id)	
	return resp.Port, nil
}

func ReceiveTarGzFromServer(id string) (string, error) {
	fullURL := fmt.Sprintf("%s/%s", BackendURL, id)
	response, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	infoURL := fmt.Sprintf("%s/info/%s", BackendURL, id)
	infoResponse, err := http.Get(infoURL)
	if err != nil {
		return "", err
	}
	defer infoResponse.Body.Close()

	port, err := GetServerInfo(id)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return port, nil
}
