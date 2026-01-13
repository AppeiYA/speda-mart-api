package utils

import (
	"net/http"
	"strconv"
)

func ParsePagination(r *http.Request) (limit, offset int) {
	q := r.URL.Query()

	page := 1
	limit = 10

	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	if l := q.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			page = v
		}
	}

	if limit > 100 {
		limit = 100
	}

	offset = (page - 1) * limit
	return limit, offset
}
