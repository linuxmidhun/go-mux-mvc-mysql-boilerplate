package data

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	data "../../models"
	"github.com/gorilla/mux"
)

// New .
func New(w http.ResponseWriter, r *http.Request) {
	var d data.Data
	_ = json.NewDecoder(r.Body).Decode(&d)
	err := data.New(d.Text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		GetAll(w, r)
		//json.NewEncoder(w).Encode(d)
	}
}

// GetAll .
func GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("All data list")
	d, err := data.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(d)
	}
}

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	nID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in converting given id to int", id)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		d, err := data.Get(nID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)
		} else {
			if d.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(http.StatusNotFound)
			} else {
				json.NewEncoder(w).Encode(d)
			}
		}
	}
}

// Edit .
func Edit(w http.ResponseWriter, r *http.Request) {
	var d data.Data
	_ = json.NewDecoder(r.Body).Decode(&d)
	err := data.Edit(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		sub := data.Data{}
		sub.ID = d.ID
		sub.Text = d.Text
		json.NewEncoder(w).Encode(sub)
	}
}

// Delete .
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	nID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in converting given id to int", id)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		err := data.Delete(nID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)
		} else {
			GetAll(w, r)
		}
	}
}
