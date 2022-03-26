package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<H1>Welcome to my awesome site!</H1>")
}

func faq(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<H1>FAQ</H1>")
	fmt.Fprintln(w,
		"<ul><li>How to learn web dev with golang? <em>Watch calhoun."+
			"io courses"+
			"</em></li></ul>")

}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>To get in touch please send an email to <a href"+
		"=\"support"+
		"@karimelnaggar.lenslocked.com\">support@karimelnaggar.lenslocked."+
		"com</a>.</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<p>Page is not found</p>")
}

func main() {
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/faq", faq)
	r.GET("/contact", contact)
	r.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}
