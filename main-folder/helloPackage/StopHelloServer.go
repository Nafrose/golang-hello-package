package hellopackage

import (
	"github.com/julienschmidt/httprouter"
)

func StopHelloServer () {
	router := httprouter.New()
	router.GET("/stop", StopServer)

}