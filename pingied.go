package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/binding"
import "fmt"

// DisplayMessage is a struct for formdata from /create/text calls
type DisplayMessage struct {
	Font string `form:"font" binding:"required"`
	Text string `form:"text" binding:"required"`
}

func createText(msg DisplayMessage) string {
    return fmt.Sprintf("Hello %s\n", msg.Text)
}

func main() {
	m := martini.Classic()
	m.Post("/create/text", binding.Bind(DisplayMessage{}), createText)
	m.Run()}
