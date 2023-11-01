package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func ParseRequest(r *http.Request, dst interface{}) error {

	// Check whether the destination type is indeed a pointer to a struct.
	// This is necessary to be able to modify the actual struct that the pointer is pointing to.
	if reflect.TypeOf(dst).Kind() != reflect.Ptr || reflect.ValueOf(dst).IsNil() {
		return fmt.Errorf("destination must be a non-nil pointer to a struct")
	}

	// Decode the body into the struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dst); err != nil {
		return fmt.Errorf("error decoding request body: %w", err)
	}

	return nil
}
