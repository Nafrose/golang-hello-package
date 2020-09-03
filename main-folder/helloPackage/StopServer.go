package hellopackage

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StopServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hold your horses!")
}