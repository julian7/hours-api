package handlers

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

// LoggingHandler logs http queries
func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := statusWriter{w, 0, 0}
		raddr, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal(err)
			return
		}
		h.ServeHTTP(&writer, r)
		log.Println(fmt.Sprintf(
			"%s %d %d \"%s %s %s\" %s",
			raddr, writer.status, writer.length,
			r.Method, r.URL.String(), r.Proto, r.UserAgent()))
	})
}
