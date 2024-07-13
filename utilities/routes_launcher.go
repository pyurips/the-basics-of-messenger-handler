package utilities

import (
	"fmt"
	"net/http"
)

func InitializeRoutes() {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/hello": h1Handler,
		"/":      notFoundHandler,
	}

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
}

func h1Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>H1</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "<h1>Page not found</h1>")
}
