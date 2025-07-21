package api

import (
	"encoding/json"
	"fmt"
	"learning_GO/RAMEN_API_base/internal/app/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status:code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArcticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get all Articles GET /api/v1/arcticles")

	articles, err := api.storage.Article().SelectAll()

	if err != nil {
		api.logger.Info("Error while Articles.SelectAll :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trouble whith accesing to database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArcticalById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get arcticle by ID api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles with parsing [id] param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id param, dont use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArcicleById(id)
	if err != nil {
		api.logger.Info("Troubles while accesing database table (arcticles) with id  err :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Problem with database, try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("arcticle with this id, does not found", err)
		msg := Message{
			StatusCode: 404,
			Message:    "Invalid id param, not arcticle with this id in DataBase",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArcticalById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete arcticle by ID api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles with parsing [id] param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id param, dont use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindArcicleById(id)
	if err != nil {
		api.logger.Info("Troubles while accesing database table (arcticles) with id  err :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Problem with database, try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("arcticle with this id, does not found", err)
		msg := Message{
			StatusCode: 404,
			Message:    "Invalid id param, not arcticle with this id in DataBase",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, err := api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles with deleting arcticle with param", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Problem with database, try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) PostArcticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("POST Article POST /api/v1/articles")
	var arcticle models.Article
	err := json.NewDecoder(req.Body).Decode(&arcticle)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&arcticle)
	if err != nil {
		api.logger.Info("Error while Articles.Create :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trouble whith accesing to database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(a)

}

func (api *API) PostUserRegistr(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("POST user Registr POST /api/v1/user/registr")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FingByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accesing database table (arcticles) with id  err :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Problem with database, try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User with that ID already exist in database")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that ID already exist in database",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
	}

	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles with deleting arcticle with param", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Problem with database, try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} succesfully registr", userAdded.Login),
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
