package main

import "github.com/raumzeitlabor/pingied/renderer"
import "github.com/raumzeitlabor/pingied/util"

import "github.com/go-martini/martini"
import "github.com/martini-contrib/binding"
import "fmt"
import "net/http"
import "mime/multipart"
import "io/ioutil"

// DisplayMessage is a struct for formdata from /create/text calls
type DisplayMessage struct {
	Font string `form:"font" binding:"required"`
	Text string `form:"text" binding:"required"`
}

// ImageMessage is a struct to contain a file from /create/image calls
type ImageMessage struct {
    Title      string                `form:"title"`
    TextUpload *multipart.FileHeader `form:"txtUpload"`
}

// IDMessage is a struct to contain a ID from /show/image and /show/scroll calls
type IDMessage struct {
	ID string `form:"id" binding:"required"`
}

func createText(msg DisplayMessage) (int, string) {
	image, err := renderer.RenderImage(msg.Text, msg.Font)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	util.StoreImage(image)
	return http.StatusOK, fmt.Sprintf("Hello %s\n", msg.Text)
}

func createImage(msg ImageMessage) (int, string) {
    file, err := msg.TextUpload.Open()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
    data, err := ioutil.ReadAll(file)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
    sha, _, err := util.StoreImage(data)
	return http.StatusNotImplemented, sha
}

func displayImage(msg IDMessage) (int, string) {
    _, err := util.RetrieveImage(msg.ID)
    if err != nil {
        return http.StatusInternalServerError, err.Error()
    }
	return http.StatusNotImplemented, msg.ID
}

func main() {
	m := martini.Classic()
	m.Post("/text", binding.Bind(DisplayMessage{}), createText)
	m.Post("/image", binding.MultipartForm(ImageMessage{}), createImage)
	m.Post("/display", binding.Bind(IDMessage{}), displayImage)
	m.Run()
}
