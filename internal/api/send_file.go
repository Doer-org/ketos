package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// type Response struct {
//     PullID string `json:"id"` // pullIDからPullIDへ変更して公開
// }

// TODO: 取り敢えず書いてみただけなので、動作確認しつつ適宜変更してください。
func SendTarToServer(publishList []string, envList []string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("upload_file", filepath.Base(filePath))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	portsString := strings.Join(publishList, ",")
	if portsString == "" {
		portsString = "none"
	}

	fullURL := fmt.Sprintf("%s?port=%s", BackendURL, portsString)
	request, err := http.NewRequest("POST", fullURL, body)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	
	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody.String()))
	var resp Response
	err = json.Unmarshal(responseBody.Bytes(), &resp)
	if err != nil {
		return err
	}
	fmt.Printf("\n\n\nShare this command!!!!\n")
	fmt.Printf("ketos pull -i %s\n", resp.ID) // resp.pullIDからresp.PullIDへ変更
	return nil
}
