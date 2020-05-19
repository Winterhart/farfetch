package farfetch

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//Farfetch - Communicate with a slack channel using a Slack-app
type Farfetch interface {
	SendMessage(message string) error
	UploadFile(pathToFile string) error
}

//NewFarfetchImpl - Constructor that implements the Farfetch API
func NewFarfetchImpl(hook string, token string, channel string) Farfetch {
	return &farfetchImpl{
		hook:    hook,
		token:   token,
		channel: channel,
	}
}

type farfetchImpl struct {
	hook, token, channel string
}

//SendMessage - This function sends a message to Slack using a web-token.
func (f *farfetchImpl) SendMessage(message string) (err error) {
	slackBodyJSON := "{\"text\":\"%v\"}"
	jsonMessage := fmt.Sprintf(slackBodyJSON, message)
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest(http.MethodPost, f.hook, bytes.NewBuffer([]byte(jsonMessage)))
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return err
}

//UploadFile - This function upload a file to Slack using a token and a channel id
func (f *farfetchImpl) UploadFile(pathToFile string) (err error) {
	fileName := filepath.Base(pathToFile)
	url := fmt.Sprintf(
		"https://slack.com/api/files.upload?token=%v&filename=%v&channels=%v&pretty=1",
		f.token,
		fileName,
		f.channel,
	)
	httpClient := &http.Client{
		Timeout: time.Minute * 10,
	}
	body, contentType, err := generateFileForm(pathToFile)
	request, err := http.NewRequest(http.MethodPost, url, body)
	request.Header.Add("Content-Type", contentType)
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return err
}

//generateFileForm - This helper function creates a web form with the desired file
func generateFileForm(pathToFile string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	p, err := filepath.Abs(pathToFile)
	if err != nil {
		return nil, "", err
	}
	file, err := os.Open(p)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	part, err := writer.CreateFormFile("file",
		filepath.Base(pathToFile))
	if err != nil {
		return nil, "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
