package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pmutua/health-insurance-api/app/models"
)

func GetAllPersons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	persons := []models.Person{}
	db.Find(&persons)
	respondJSON(w, http.StatusOK, persons)
}

func CreatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	person := models.Person{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, person)
}

func GetPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func UpdatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func DeletePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	if err := db.Delete(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

//getPersonOr404 gets person instance if exists , or respond the 404 error or otherwise
func getPersonOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *models.Person {
	person := models.Person{}
	if err := db.First(&person, models.Person{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &person
}
