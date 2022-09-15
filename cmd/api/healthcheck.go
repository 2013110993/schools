//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status": "available", "environment": %q, "version": %q }`
	//js = fmt.Sprintf(js, app.config.env, version)

	//Create a map to hold our healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	//Convert our map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encontered a problem and could process your request", http.StatusInternalServerError)
		return
	}
	//Add a newline to make viewing on the terminal easier
	js = append(js, '\n')
	//Specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")

	//Write the []byte slice containing the JSON response body
	w.Write(js)

	//Write the JSON as the HTTP response body
	//w.Write([]byte(js))

	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
}
