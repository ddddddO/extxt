package extxt

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Extxt</title>
</head>
<body>
	<form action="/" method="post" enctype="multipart/form-data">
		<input type="file" name="src_file">
		<button type="submit">text!</button>
	</form>
</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if !basicAuthenticated(r) {
		w.Header().Add("WWW-Authenticate", `Basic realm="secret xxxx"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprint(w, indexHTML)
		return
	}

	if r.Method == http.MethodPost {
		f, header, err := r.FormFile("src_file")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		_ = header

		if err := RunByServer(w, f); err != nil {
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
