package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/m00nreal/go-jwt-auth/api/database"
	"github.com/m00nreal/go-jwt-auth/api/models"
	"github.com/m00nreal/go-jwt-auth/api/repository"
	"github.com/m00nreal/go-jwt-auth/api/repository/crud"
	"github.com/m00nreal/go-jwt-auth/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCrud(db)

	func (usersRepository repository.UserRepository) {
		users, err := usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	repo := crud.NewRepositoryUsersCrud(db)

	func (usersRepository repository.UserRepository) {
		user, err := usersRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCrud(db)

	func (usersRepository repository.UserRepository) {
		user, err := usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	repo := crud.NewRepositoryUsersCrud(db)

	func (usersRepository repository.UserRepository) {
		_, err := usersRepository.Delete(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	repo := crud.NewRepositoryUsersCrud(db)

	func (usersRepository repository.UserRepository) {
		rows, err := usersRepository.Update(uint32(uid), user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}