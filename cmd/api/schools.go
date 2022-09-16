package main

import (
	"fmt"
	"net/http"
	"time"

	"schools.federicorosado.net/internal/data"
)

//createSchoolHandler for the "POST" /v1/schools  endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new school..")
}

//showSchoolHandler for the "GET" /v1/schools/:id  endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Create a new instance of the School struct containing the ID we extracted
	//from our URL and some sample data
	school := data.School{
		ID:       id,
		CreateAt: time.Now(),
		Name:     "Apple Tree",
		Level:    "High School",
		Contact:  "Anna Smith",
		Phone:    "601-4411",
		Address:  "14 Apple Street",
		Mode:     []string{"blended", "online"},
		Version:  1,
	}
	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a proble and could not process your request", http.StatusInternalServerError)
	}
}
