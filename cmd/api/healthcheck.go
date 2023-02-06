package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status:Avialable")
	fmt.Fprintf(w, "env %s\n", app.config.env)
	fmt.Fprintf(w, "version %s", version)
}
