package main

import (
	hellopackage "main-folder/main-folder/helloPackage"
)

func main (){
	h := hellopackage.HttpHelper{}
	h.StartHelloServer()
	h.StopHelloServer()
}