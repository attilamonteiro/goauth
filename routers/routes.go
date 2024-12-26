package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"goauth/controllers"
	"goauth/core"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	// Ajuste 1: Utilizar negroni.Wrap para funções padrão (http.HandlerFunc)
	router.Handle("/refresh-token",
		negroni.New(
			negroni.HandlerFunc(core.RequireTokenAuthentication),
			negroni.Wrap(http.HandlerFunc(controllers.RefreshToken)),
		)).Methods("GET")

	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(core.RequireTokenAuthentication),
			negroni.Wrap(http.HandlerFunc(controllers.Logout)),
		)).Methods("POST")

	return router
}
