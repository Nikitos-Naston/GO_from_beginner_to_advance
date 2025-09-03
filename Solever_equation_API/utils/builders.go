package utils

import (
	handlers "Solver_equation_API/Handlers"

	"github.com/gorilla/mux"
)

func BuildSolverResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/solve", handlers.SolveEquation).Methods("GET")
	router.HandleFunc(prefix+"/grab", handlers.GrabKoef).Methods("POST")

}
