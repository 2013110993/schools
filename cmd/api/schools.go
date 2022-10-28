package main

import (
	"errors"
	"fmt"
	"net/http"

	"schools.federicorosado.net/internal/data"
	"schools.federicorosado.net/internal/validator"
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
	//Copy the values from the input struct to a new schools struct
	school := &data.School{
		Name:    input.Name,
		Level:   input.Level,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode:    input.Mode,
	}

	//Initialize a new validator instance
	v := validator.New()

	//Check the map to determin if there were any validation errors
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Create a school
	err = app.models.Schools.Insert(school)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// create a location header for the newly created resource/school
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/schools/%d", school.ID))
	// write the JSON response with 201 - Created status code with the body
	//being the school data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"school": school}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//showSchoolHandler for the "GET" /v1/schools/:id  endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Fetch the specifi school
	school, err := app.models.Schools.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// //Create a STATIC new instance of the School struct containing the ID we extracted
	// //from our URL and some sample data
	// school := data.School{
	// 	ID:        id,
	// 	CreatedAt: time.Now(),
	// 	Name:      "Apple Tree",
	// 	Level:     "High School",
	// 	Contact:   "Anna Smith",
	// 	Phone:     "601-4411",
	// 	Address:   "14 Apple Street",
	// 	Mode:      []string{"blended", "online"},
	// 	Version:   1,
	// }

	//Write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Create Update Method
func (app *application) updateSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//This method does a partial replacement
	//Get the id for the school that needs updating
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//Fetch the original record from the database
	school, err := app.models.Schools.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Create an input struct to hold data read in from the client
	// we update the input struct to use pointers because pointers have a
	// default value of nill
	//If the field remains nil then we know that the client did not update it
	var input struct {
		Name    *string  `json:"name"`
		Level   *string  `json:"level"`
		Contact *string  `json:"contact"`
		Phone   *string  `json:"phone"`
		Email   *string  `json:"email"`
		Website *string  `json:"website"`
		Address *string  `json:"address"`
		Mode    []string `json:"mode"`
	}

	//Initialize a new Json.Decoder instance
	err = app.readJSON(w, r, &input) //json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Check for updates
	if input.Name != nil {
		school.Name = *input.Name
	}
	if input.Level != nil {
		school.Level = *input.Level
	}
	if input.Contact != nil {
		school.Contact = *input.Contact
	}
	if input.Phone != nil {
		school.Phone = *input.Phone
	}
	if input.Email != nil {
		school.Email = *input.Email
	}
	if input.Website != nil {
		school.Website = *input.Website
	}
	if input.Address != nil {
		school.Address = *input.Address
	}
	if input.Mode != nil {
		school.Mode = input.Mode
	}
	// Copy/update the fields/values in the school variable using the fields
	// in the input struct
	// school.Name = input.Name
	// school.Level = input.Level
	// school.Contact = input.Contact
	// school.Phone = input.Phone
	// school.Email = input.Email
	// school.Website = input.Website
	// school.Address = input.Address
	// school.Mode = input.Mode

	// Perform validation on the update school. If validation fails, then
	// we send a 422 - unprocessable Entity response to the client

	//Initialize a new validator instance
	v := validator.New()

	//Check the map to determin if there were any validation errors
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the update school record to the update method
	err = app.models.Schools.Update(school)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//Write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//Delete Handler method
func (app *application) deleteSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//Get the id for the school that needs updating
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the School from the database. Send a 404 Not Found Status code to the
	//cliet if there is no matching record
	err = app.models.Schools.Delete(id)

	//Handle erros
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return 200 status OK to the client with a success message
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "school successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// The listSchoolsHandler() allows the client to see a listing of schools
// based on a set of criteria
func (app *application) listSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Create an input struct to hold our query parameters
	var input struct {
		Name  string
		Level string
		Mode  []string
		data.Filters
	}
	// Initialize a validator
	v := validator.New()

	// Get the URL values map
	qs := r.URL.Query()

	//Use the helper method to extract the values
	input.Name = app.readString(qs, "name", "")
	input.Level = app.readString(qs, "level", "")
	input.Mode = app.readCSV(qs, "mode", []string{})
	// Get the informaton
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	// Get the sort information
	input.Filters.Sort = app.readString(qs, "sort", "id")
	// Specific the allowed sort values
	input.Filters.SortList = []string{"id", "name", "level", "-id", "-name", "-level"}
	// Check for validation errors
	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// //Results Dump
	// fmt.Fprintf(w, "%+v\n", input)

	// Get a listing of all schools
	schools, err := app.models.Schools.GetAll(input.Name, input.Level, input.Mode, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response conting all response
	err = app.writeJSON(w, http.StatusOK, envelope{"schools": schools}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
