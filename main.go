package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./mail"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mail.InitConfig()

	http.HandleFunc("/", idx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Running over localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Interno - Contactar al administrador", http.StatusInternalServerError)
	}

}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Interno - Contactar al administrador", http.StatusInternalServerError)
	}

}

func blog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "blog.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Interno - Contactar al administrador", http.StatusInternalServerError)
	}

}

func contact(w http.ResponseWriter, req *http.Request) {
	var correo mail.Mail

	if req.Method == http.MethodPost {
		correo = mail.Mail{
			Nombre:   req.FormValue("nombre"),
			Correo:   req.FormValue("correo"),
			Telefono: req.FormValue("telefono"),
			Mensaje:  req.FormValue("mensaje"),
		}

		mail.Send(correo)
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Interno - Contactar al administrador", http.StatusInternalServerError)
	}

}
