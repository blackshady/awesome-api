package middlewares

import (
	"net/http"
)

func LoggingMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//	log any the request information or perform middleware action
	next(res, req)
}
