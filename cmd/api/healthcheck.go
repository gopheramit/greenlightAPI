package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "Available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJson(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		// app.logger.Println(err)
		// http.Error(w, "server encountered error and could not process your request", http.StatusInternalServerError)
		// return
		app.serverErrorResponse(w, r, err)
	}

}
