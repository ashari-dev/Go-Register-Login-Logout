package dtos

type Respont struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}
