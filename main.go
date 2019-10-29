package main

import (
	"bytes"
	"fmt"
	"text/template"
	"log"
	"net/http"
	"os"
	"strings"
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
	Headers map[string]string
	Cookies map[string]string
	PostVar string
	GetVar string
}

func (r RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Read headers into map
	var data = TemplateData{}
	data.Headers = make(map[string]string)
	for k, v := range req.Header {
		data.Headers[k] = strings.Join(v, ",")
	}

	// Read cookies into map
	data.Cookies = make(map[string]string)
	for _, cookie := range req.Cookies() {
		data.Cookies[cookie.Name] = cookie.Value
	}

	// Read POST variables
	err := req.ParseForm()
	if err != nil {
		internalError(err, w)
	}
	data.PostVar = req.Form.Get("formvar")
	data.GetVar = req.URL.Query().Get("formvar")

	// Render Response
	t := template.Must(template.New("index.html").ParseFiles("public/index.html"))
	buf := bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "index", data)
	if err != nil {
		internalError(err, w)
	}

	w.Write(buf.Bytes())

}

func internalError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
