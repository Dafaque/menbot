package webhandler

import (
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path"
)

type handler struct {
	fs          fs.FS
	prefix      string
	defaultFile string
}

func init() {
	mime.AddExtensionType(".ts", "text/javascript")
}

func New(fs fs.FS, prefix string, defaultFile string) http.Handler {
	return &handler{fs: fs, prefix: prefix, defaultFile: defaultFile}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file string = h.prefix
	if r.URL.Path == "/" {
		file = path.Join(h.prefix, h.defaultFile)
	} else {
		file = path.Join(h.prefix, r.URL.Path)
	}

	ext := path.Ext(file)
	mtype := mime.TypeByExtension(ext)
	if mtype != "" {
		w.Header().Set("Content-Type", mtype)
	}

	f, err := h.fs.Open(file)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
