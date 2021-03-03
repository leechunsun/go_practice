package requestAndResponse


type UpperCaseRequest struct {
	S string `json:"s"`
}

type UpperCaseResponse struct {
	V string `json:"v"`
	E string `json:"e,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

