package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<H1>Welcome to my awesome site!</H1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"support"+
			"@karimelnaggar.lenslocked.com\">support@karimelnaggar.lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found!")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
