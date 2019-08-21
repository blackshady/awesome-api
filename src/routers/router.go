package routers

import (
	"awesome-api/src/controllers"
	"awesome-api/src/middlewares"
	"awesome-api/src/urls"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"net/http"
)

func GetRouter() *pat.Router {
	url_patterns := urls.ReturnURLS()

	server := pat.New()
	server.Get(url_patterns.HOME_PATH, controllers.HomeController)

	common := pat.New()
	common.PathPrefix(url_patterns.STATIC_PATH).Handler(http.StripPrefix(url_patterns.STATIC_PATH,
		http.FileServer(http.Dir("static"))))

	common.PathPrefix(url_patterns.PROJECT_VERSION_PATH).Handler(
		negroni.New(negroni.HandlerFunc(
			middlewares.LoggingMiddleware), negroni.Wrap(server)),
	)
	return common
}
