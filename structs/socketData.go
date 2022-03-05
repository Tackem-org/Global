package structs

type SocketRequest struct {
	Command string
	User    *UserData
	Data    map[string]interface{}
}

type SocketReturn struct {
	StatusCode   uint32
	ErrorMessage string
	TellAll      bool
	Data         map[string]interface{}
}
