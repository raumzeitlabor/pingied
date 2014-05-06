package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/binding"
import "fmt"
import "net/http"

// DisplayMessage is a struct for formdata from /create/text calls
type DisplayMessage struct {
	Font string `form:"font" binding:"required"`
	Text string `form:"text" binding:"required"`
}

// ImageMessage is a struct to contain a file from /create/image calls
type ImageMessage struct {
}

// IDMessage is a struct to contain a ID from /show/image and /show/scroll calls
type IDMessage struct {
    ID string `form:"id" binding:"required"`
}

func createText(msg DisplayMessage) (int, string) {
    return http.StatusOK, fmt.Sprintf("Hello %s\n", msg.Text)
}

func createImage(_ ImageMessage) (int, string) {
    var sha = ""
    return http.StatusNotImplemented, sha
}

func showImage(_ IDMessage) (int, string) {
    var sha = ""
    return http.StatusNotImplemented, sha
}

func main() {
	m := martini.Classic()
	m.Post("/create/text", binding.Bind(DisplayMessage{}), createText)
	m.Post("/create/image", binding.Bind(ImageMessage{}), createImage)
	m.Post("/show", binding.Bind(IDMessage{}), showImage)
	m.Run()
}
