package controllers

import "net/http"

func Post(f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f(w, r)
	}
}
