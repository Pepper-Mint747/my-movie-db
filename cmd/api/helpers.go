package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// readIDParam retrieves the "id" URL parameter frfom the current request context,
//then converts it to an integer and return it
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define  an envelope type
type envelope map[string]interface{}

// writeJSON takes the destination http.ResponseWriter, the HTTP
//status code to send, the data to encode to JSON, and a header
//map containing any additional HTTP headers we want to include in the response

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications
	js = append(js, '\n')

	// We loop through the header map and add each header to the http.ResponseWriter header map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	//Add the "Content-type header then write the code and JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil

}
