package controller

import (
	"fmt"
	"net/http"
)

func HandlerHomepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello internet")
}
