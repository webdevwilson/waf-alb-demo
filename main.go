package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
)

var url, wafUrl, port string

type RequestHandler struct{}

func main() {

	port = envVar("PORT", "80")

	defaultUrl := fmt.Sprintf("http://localhost:%s/", port)
	wafUrl = envVar("WAF_URL", defaultUrl)
	url = envVar("URL", defaultUrl)

	r := RequestHandler{}
	http.Handle("/", r)

	log.Printf("[INFO] Starting HTTP server on port %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func envVar(key string, defaultVal string) string {
	var val string
	var ok bool
	if val, ok = os.LookupEnv(key); !ok {
		val = defaultVal
	}
	return val
}

type TemplateData struct {
	HasPayload  bool
	Payload     string
	PayloadHTML template.HTML
	Url         string
	WafUrl      string
}

func (r RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Create template data struct
	var data = TemplateData{
		Url:    url,
		WafUrl: wafUrl,
	}

	// On post, grab the form value
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
