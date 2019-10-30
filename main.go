package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
)

type RequestHandler struct {}

func main() {

	var port string
	var ok bool
	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "80"
	}

	r := RequestHandler{}
	http.Handle("/", r)

	log.Printf("[INFO] Starting HTTP server on port %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

type TemplateData struct {
	HasPayload bool
	Payload string
	PayloadHTML template.HTML
}

func (r RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Read headers into map
	var data = TemplateData{}

	if req.Method == "POST" {
		req.ParseForm()
		data.Payload = req.Form.Get("payload")
	}

	data.PayloadHTML = template.HTML(data.Payload)
	data.HasPayload = data.Payload != ""

	// Render Response
	t := template.Must(template.New("index.html").ParseFiles("public/index.html"))
	buf := bytes.NewBuffer([]byte{})
	err := t.ExecuteTemplate(buf, "index", data)
	if err != nil {
		internalError(err, w)
	}

	w.Write(buf.Bytes())

}

func internalError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
