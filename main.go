package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my go website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Contact Page</h1><p>get in touch with my
	 email: <a href="mailto:behzad@engineer.com">behzad@engineer.com</a>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to faq page!</h1>")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "The chosen name is: %s", name)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/names/{name}", nameHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!", http.StatusNotFound)
	})
	fmt.Println("Staring the server on port 3000...")
	http.ListenAndServe(":3000", r)
}
