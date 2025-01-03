package structs

type PanelRequest struct {
	Name      string
	User      *UserData
	Variables map[string]string
}

type PanelReturn struct {
	StatusCode   uint32
	ErrorMessage string
	HTML         string
}
