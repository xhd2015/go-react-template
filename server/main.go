package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := handle()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func handle() error {
	staticDir := "./dist"
	fs := http.FileServer(http.Dir(staticDir))
	var serveFiles http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: Method=%s, Host=%s, RemoteAddr=%s, RequestURI=%s, Header=%v", r.Method, r.Host, r.RemoteAddr, r.RequestURI, r.Header)

		file := filepath.Join(staticDir, strings.ReplaceAll(r.RequestURI, "/", string(filepath.Separator)))
		fsInfo, statErr := os.Stat(file)
		if statErr != nil {
			if !errors.Is(statErr, os.ErrNotExist) {
				http.Error(w, statErr.Error(), 500)
				return
			}
		} else if !fsInfo.IsDir() {
			// is file
			osFile, err := os.Open(file)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer osFile.Close()
			http.ServeContent(w, r, file, fsInfo.ModTime(), osFile)
			return
		}
		fs.ServeHTTP(w, r)
	}
	http.Handle("/", serveFiles)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%s...", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
