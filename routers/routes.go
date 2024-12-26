package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"goauth/controllers"
	"goauth/repository"
	"net/http"
	"goauth/core"
)

func SetAuthenticationRoutes(router *mux.Router, userRepo repository.UserRepository) {
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(w, r, userRepo) // Passa o repositório aqui
	}).Methods("POST")

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.Register(w, r, userRepo) // Passa o repositório aqui
	}).Methods("POST")

	router.Handle("/refresh-token",
		negroni.New(
			negroni.HandlerFunc(auth.RequireTokenAuthentication),
			negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				controllers.RefreshToken(w, r, userRepo) // Passa o repositório aqui
			})),
		)).Methods("GET")

	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(auth.RequireTokenAuthentication),
			negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				controllers.Logout(w, r, userRepo) // Passa o repositório aqui
			})),
		)).Methods("POST")
}
