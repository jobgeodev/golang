package goutil

// shell
// go mod init github.com/jobgeodev/golang
// go: creating new go.mod: module github.com/jobgeodev/golang

import (
	"bytes"
	"encoding/binary"
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

// 使用BigEndian方式将uint32转为Bytes
func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

// 使用BigEndian方式将Bytes转为uint32
func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}

// 将多个Bytes合并成一个Bytes
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte{})
}

func BuildCustomData(raw_data []byte) []byte {
	len := len(raw_data)
	len_data := Uint32ToBytes(uint32(len))

	// 新数据 = 原数据长度[uint32 四个字节] + 原数据
	custom_data := BytesCombine(len_data, raw_data)
	return custom_data
}
