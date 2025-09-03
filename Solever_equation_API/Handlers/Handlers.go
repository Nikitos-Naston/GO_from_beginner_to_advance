package handlers

import (
	helpers "Solver_equation_API/Helpers"
	"Solver_equation_API/models"
	"encoding/json"
	"log"
	"net/http"
)

var koefs models.Coef

func initHandlers(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GrabKoef(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)
	log.Println("Starting grabbing koefs")

	err := json.NewDecoder(request.Body).Decode(&koefs)

	if err != nil {
		log.Println("error is in the json file", err)
		writer.WriteHeader(401)
		json.NewEncoder(writer).Encode("error with your json file")
		return
	}

	log.Println("Koef added succesfull")
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(koefs)

}

func SolveEquation(writer http.ResponseWriter, request *http.Request) {
	log.Println("Starting calculation")
	initHandlers(writer)
	var answer models.Answer
	helpers.Convertor(&koefs, &answer)

	if answer.A != 0 {
		disk := answer.B*answer.B - 4*answer.C*answer.A
		if disk > 0 {
			answer.Roots = 2
		} else if disk == 0 {
			answer.Roots = 1
		}
	} else if answer.A == 0 && answer.B != 0 && answer.C != 0 {
		answer.Roots = 1
	}
	log.Println("End calculation")

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(answer)
}
