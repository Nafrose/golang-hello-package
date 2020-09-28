package hellopackage

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func PrintHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello")
}