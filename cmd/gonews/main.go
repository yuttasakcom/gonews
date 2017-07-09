package main

import (
	"net/http"
	"os"

	"github.com/yuttasakcom/gonews/pkg/app"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("static")})))
	http.ListenAndServe(port, mux)
}

type noDir struct {
	http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
