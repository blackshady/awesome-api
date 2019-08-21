package controllers

import (
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func HomeController(res http.ResponseWriter, req *http.Request) {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	if name := req.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(res, "welcome-template.html", welcome); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
