package structs

type WebRequest struct {
	FullPath      string
	CleanPath     string
	User          *UserData
	SessionToken  string
	Method        string
	QueryParams   map[string]interface{}
	Post          map[string]interface{}
	PathVariables map[string]interface{}
}

type WebReturn struct {
	StatusCode     uint32
	ErrorMessage   string
	FilePath       string
	PageString     string
	PageData       map[string]interface{}
	CustomPageName string
	CustomCss      []string
	CustomJs       []string
}
