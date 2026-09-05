// Фрагмент 2: book/chapters/15-adapters.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "net/http"
func serve(w http.ResponseWriter, r *http.Request) {}

var handler http.Handler = http.HandlerFunc(serve)
