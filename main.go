package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		if name := request.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(writer, "welcome-template.html", welcome); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Server started on port 8080")

	fmt.Println(http.ListenAndServe(":8080", nil))
}
