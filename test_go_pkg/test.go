package main

import (
	"log"

	goutil "github.com/jobgeodev/golang"
)

// shell
// go mod init github.com/jobgeodev/golang
// go: creating new go.mod: module github.com/jobgeodev/golang

// Hello says hello.
// func Hello() {
// 	log.Println("Hello go mod!")
// }

// // Bye says bye.
// func Bye() {
// 	log.Println("Bye go mod!")
// }

func main() {
	goutil.PostRequest("aa", "")

	data := []byte("helloworld")
	v1 := goutil.BuildCustomData(data)
	log.Printf("[%v]->[%v]", data, v1)
}
