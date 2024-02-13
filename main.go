package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/konigbach/lenslocked/controllers"
	"github.com/konigbach/lenslocked/templates"
	"github.com/konigbach/lenslocked/views"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/shop", controllers.StaticHandler(views.Must(views.ParseFs(templates.FS, "shop.gohtml", "tailwind.gohtml"))))

	var usersC controllers.Users
	usersC.Templates.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server...")

	http.ListenAndServe("localhost:3000", r)
}
