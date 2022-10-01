package httputils

import (
	"net/http"
	"strings"
)

func RemoveTrailingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

		next.ServeHTTP(w, r)
	})
}

// Appends "localhost" to address on webserver initialization when app env is "dev"
// (prevents macOS annoying Firewall request for authorization everytime the application restarts)
func GetAddr(env, port string) string {
	if env == "dev" {
		return "localhost:" + port
	}

	return ":" + port
}
