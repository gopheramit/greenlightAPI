package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {

	app.logger.Println(err)
}

func (app *application) errorRespone(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJson(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)

	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "server encountered a problem and could not process your request."
	app.errorRespone(w, r, http.StatusInternalServerError, message)

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorRespone(w, r, http.StatusNotFound, message)

}

func (app *application) methodNotFoung(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this reosurce")
	app.errorRespone(w, r, http.StatusMethodNotAllowed, message)
}