package goutil

// shell
// go mod init github.com/jobgeodev/golang
// go: creating new go.mod: module github.com/jobgeodev/golang

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Hello says hello.
func Hello() {
	log.Println("Hello go mod!")
}

// Bye says bye.
func Bye() {
	log.Println("Bye go mod!")
}

func PostRequest(url string, body_text string) ([]byte, error) {
	data, err := SendRequest(true, url, body_text)
	return data, err
}

func GetRequest(url string, body_text string) ([]byte, error) {
	data, err := SendRequest(false, url, body_text)
	return data, err
}

func SendRequest(is_post bool, url string, body_text string) ([]byte, error) {
	RequestType := "POST"
	if !is_post {
		RequestType = "GET"
	}

	log.Printf("SendRequest [%s] : [%s]", RequestType, url)

	request, err := http.NewRequest(RequestType, url, strings.NewReader(body_text))
	if err != nil {
		log.Printf("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		b = nil
		log.Printf("http.Do failed,[StatusCode=%d] [err=%s][url=%s]", resp.StatusCode, err, url)
	}
	return b, err
}
