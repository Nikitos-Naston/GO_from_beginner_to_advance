package main

import (
	"Solver_equation_API/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port      string = "8100"
	apiPrefix string = "/api"
)

func main() {
	log.Println("Trying to start Solver Rest API server")

	router := mux.NewRouter()

	log.Println("Initialization Succes")

	utils.BuildSolverResource(router, apiPrefix)
	log.Println("starting lissing port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
