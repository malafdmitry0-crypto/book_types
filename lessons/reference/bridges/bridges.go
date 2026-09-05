package bridges

import (
	"errors"
	"io/fs"
	"net/http"
)

// imports: "errors", "io/fs"
func pathFromError(err error) (string, bool) {
	var target *fs.PathError
	if errors.As(err, &target) {
		return target.Path, true
	}
	return "", false
}

// import "net/http"
func serve(w http.ResponseWriter, r *http.Request) {}

var handler http.Handler = http.HandlerFunc(serve)
