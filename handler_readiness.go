package main

import "net/http"

func handlerReadniess(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
