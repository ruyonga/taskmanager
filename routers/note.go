package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/ruyonga/taskmanager/common"
	"github.com/ruyonga/taskmanager/controllers"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noterRouter := mux.NewRouter()

	noterRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noterRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noterRouter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
	noterRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noterRouter.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	noterRouter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noterRouter),
	))
	return router

}
