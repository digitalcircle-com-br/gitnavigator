//go:build !dev

package static

import (
	"embed"
	"mime"
	"net/http"
	"path"
)

//go:embed webroot
var fsroot embed.FS

func Serve(w http.ResponseWriter, r *http.Request) {
	pt := r.URL.Path
	if pt == "" || pt == "/" {
		pt = "/index.html"
	}
	pt = pt[1:]
	bs, err := fsroot.ReadFile(path.Join("webroot", pt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mt := mime.TypeByExtension(path.Ext(pt))
	w.Header().Add("Content-Type", mt)
	w.Write(bs)
}
