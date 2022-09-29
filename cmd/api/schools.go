package main

import (
	"fmt"
	"net/http"
	"time"

	"schools.federicorosado.net/internal/data"
)

//createSchoolHandler for the "POST" /v1/schools  endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	//Initialize a new Json.Decoder instance
	err := app.readJSON(w, r, &input) //json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//Display the reques
	fmt.Fprintf(w, "%+v\n", input)
}

//showSchoolHandler for the "GET" /v1/schools/:id  endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
