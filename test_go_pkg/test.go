package main

import (
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
}
