//go:build dev

package static

import (
	"mime"
	"net/http"
	"os"
	"path"
)

func init() {
	println("Using static dev")
}
func Serve(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	pt := r.URL.Path
	if pt == "" || pt == "/" {
		pt = "/index.html"
	}
	pt = pt[1:]
	bs, err := os.ReadFile(path.Join(wd, "lib/static/webroot", pt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mt := mime.TypeByExtension(path.Ext(pt))
	w.Header().Add("Content-Type", mt)
	w.Write(bs)
}
