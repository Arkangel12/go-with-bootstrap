package mail

import (
	"os"
	"log"
	"encoding/json"
	"net/smtp"
)

type configuration struct {
	Email string `json: "email"`
	Password string `json: "password"`
	Oemail string `json: "oemail"`
}

type EnviarCorreo struct {
	Nombre   string
	Correo   string
	Telefono string
	Mensaje  string
}

var AppConfig configuration

func InitConfig()  {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("data/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}


func Send(correo EnviarCorreo) {
	from := AppConfig.Email
	pass := AppConfig.Password
	to := AppConfig.Oemail

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
