package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type HTMLDir struct {
	D http.Dir
}

func FileHandler(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			notFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func (d HTMLDir) Open(name string) (http.File, error) {
	// Try name with added extension
	f, err := d.D.Open(name + ".html")
	if os.IsNotExist(err) {
		// Not found, try again with name as supplied.
		if f, err := d.D.Open(name); err == nil {
			return f, nil
		}
	}
	return f, err
}

func notFound(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("public/404.html")

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
	}

	w.Write([]byte(f))
}
