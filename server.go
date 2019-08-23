package main

import (
	"awesome-api/src/routers"
	"github.com/urfave/negroni"
	"log"
)

func main() {
	router := routers.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening")
	n.Run(":8080")
}
