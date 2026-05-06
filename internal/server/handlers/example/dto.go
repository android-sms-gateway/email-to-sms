package example

type Request struct {
	Value string `json:"value" validate:"required"`
}

type Response struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}
