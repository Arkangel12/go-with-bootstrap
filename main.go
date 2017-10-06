package main

import (
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	//"github.com/Arkangel12/go-with-bootstrap/common"
	"./common"
)

type enviarCorreo struct {
	Nombre   string
	Correo   string
	Telefono string
	Mensaje  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	common.InitConfig()

	http.HandleFunc("/", idx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
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
	var correo enviarCorreo

	if req.Method == http.MethodPost {
		correo = enviarCorreo{
			Nombre:   req.FormValue("nombre"),
			Correo:   req.FormValue("correo"),
			Telefono: req.FormValue("telefono"),
			Mensaje:  req.FormValue("mensaje"),
		}

		send(correo)
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Interno - Contactar al administrador", http.StatusInternalServerError)
	}

}

func send(correo enviarCorreo) {
	from := common.AppConfig.Email
	pass := common.AppConfig.Password
	//to := correo.Correo
	to := common.AppConfig.Oemail

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		correo.Mensaje

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))


	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Your message has been sent!")
}
