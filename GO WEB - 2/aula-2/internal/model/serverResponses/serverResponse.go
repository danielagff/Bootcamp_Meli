package serverresponses

type ServerResponse struct {
	Data interface{} `json:"data,omitempty"`
	Message string   `json:"message"`
}