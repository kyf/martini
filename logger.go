package martini

import (
	"github.com/kyf/util/log"
	"net/http"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() Handler {
	return func(res http.ResponseWriter, req *http.Request, c Context, logger *log.Logger) {
		method, path, ip, userAgent := req.Method, req.URL.Path, getRealIP(req), req.UserAgent()

		c.Next()

		rw := res.(ResponseWriter)
		code := rw.Status()

		logger.Printf("[%s]%s  %s  %d  %s", method, ip, path, code, userAgent)
	}
}

func getRealIP(req *http.Request) string {
	addr := req.Header.Get("X-Real-IP")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
		}
	}

	return addr
}
