package param

type Request struct {
	Question string `json:"question"`
}

type Response struct {
	Request
	Answer string `json:"answer"`
}
