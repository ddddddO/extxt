package server

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ddddddO/extxt"
	tmpl "github.com/ddddddO/extxt/server/templates"
)

// RunServer is ...
func RunServer() error {
	log.Println("start")

	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}

const src = "src_file"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if !basicAuthenticated(r) {
		w.Header().Add("WWW-Authenticate", `Basic realm="secret xxxx"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		t, err := template.New("index").Parse(tmpl.IndexHTML)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err := t.Execute(w, src); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		f, header, err := r.FormFile(src)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		_ = header

		buf := &bytes.Buffer{}
		if err := extxt.RunByServer(buf, f); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		tmp := &struct {
			Text  string   `json:"Text"`
			Words []string `json:"Words"`
		}{}

		if err := json.Unmarshal(buf.Bytes(), tmp); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		t, err := template.New("extxt").Parse(tmpl.ExtxtHTML)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err := t.Execute(w, tmp); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

var (
	validNames     = strings.Split(os.Getenv("BASIC_AUTH_NAMES"), ",")
	validPasswords = strings.Split(os.Getenv("BASIC_AUTH_PASSWORDS"), ",")
)

func basicAuthenticated(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	for i := range validNames {
		if username != validNames[i] {
			continue
		}
		if password == validPasswords[i] {
			return true
		}
	}
	return false
}
