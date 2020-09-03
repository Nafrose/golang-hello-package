package hellopackage

type HttpHelper struct {

}

func (h *HttpHelper) StartHelloServer() {
	StartHelloServer()
}

func (h *HttpHelper) StopHelloServer (){
	StopHelloServer()
}