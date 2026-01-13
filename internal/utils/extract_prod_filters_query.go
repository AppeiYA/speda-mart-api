package utils

import (
	"net/http"
	"strconv"
)

func ExtractProductFilters(r *http.Request) (
	name, color, origin *string,
	minPrice, maxPrice *int64,
	err error,
) {
	q := r.URL.Query()

	if v := q.Get("name"); v != "" {
		val := v
		name = &val
	}

	if v := q.Get("color"); v != "" {
		val := v
		color = &val
	}

	if v := q.Get("origin"); v != "" {
		val := v
		origin = &val
	}

	if v := q.Get("min_price"); v != "" {
		p, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		minPrice = &p
	}

	if v := q.Get("max_price"); v != "" {
		p, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		maxPrice = &p
	}

	return name, color, origin, minPrice, maxPrice, err
}
