package secular

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

var requestBody bytes.Buffer

func Upload(body []byte, fileName string) (string, error) {

	url := "http://localhost:8080/upload"
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err

	}

	_, err = part.Write(body)
	if err != nil {
		return "", err

	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return "", err

	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err

	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}

	if strings.Contains(string(responseBody), "error") {
		return "", err
	}

	var link map[string]string
	err = json.Unmarshal(responseBody, &link)
	if err != nil {
		return "", err
	}

	return url + "s/" + link["link"], nil
}
