package server

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ddddddO/extxt"
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
		t := template.Must(template.ParseFiles("server/templates/index.html"))
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

		if err := extxt.RunByServer(w, f); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

var (
	validName     = os.Getenv("BASIC_AUTH_NAME")
	validPassword = os.Getenv("BASIC_AUTH_PASSWORD")
)

func basicAuthenticated(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	if username == validName && password == validPassword {
		return true
	}
	return false
}
