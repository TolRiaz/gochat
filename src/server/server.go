package server

import (
    "io/ioutil"
    "net/http"
    "path/filepath"
	"log"
	"fmt"
)

func Start() {
    http.Handle("/", new(staticHandler))

    err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(8080)")
	}
}

type staticHandler struct {
    http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    localPath := "html" + req.URL.Path
    content, err := ioutil.ReadFile(localPath)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte(http.StatusText(404)))
        return
    }

    contentType := getContentType(localPath)
    w.Header().Add("Content-Type", contentType)
    w.Write(content)
}

func getContentType(localPath string) string {
    var contentType string
    ext := filepath.Ext(localPath)

    switch ext {
    case ".html":
        contentType = "text/html"
    case ".css":
        contentType = "text/css"
    case ".js":
        contentType = "application/javascript"
    case ".png":
        contentType = "image/png"
    case ".jpg":
        contentType = "image/jpeg"
    default:
        contentType = "text/plain"
    }

    return contentType
}
