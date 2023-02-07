package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) readIdParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//makes data pretty for terminal but comes with cost of performance
	//js,err:=json.MarshalIndent(data)
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contain badly formatted json at %d", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contain badly formatted json")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains  incorrect  json type for field %q", unmarshalTypeError.Field)

			}
			return fmt.Errorf("body contains incorrect json, type (at character %d) ", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		case strings.HasPrefix(err.Error(), "json:unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json:unknown field")
			return fmt.Errorf("body contians unkonwn key %s", fieldName)

		case err.Error() == "http:request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		default:
			return err
		}

	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain single json value")
	}

	return nil
}
