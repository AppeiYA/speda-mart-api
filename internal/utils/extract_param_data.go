package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ExtractParams(r *http.Request, params ...string) map[string]string {
	fetched := make(map[string]string)

	pathParams := mux.Vars(r)

	for _, p := range params {
		if v, ok := pathParams[p]; ok && v != "" {
			fetched[p] = v
		}
	}

	return fetched
}