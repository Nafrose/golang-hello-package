package hellopackage

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func StartHelloServer() {
 	router := httprouter.New()
	router.GET("/", PrintHello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
