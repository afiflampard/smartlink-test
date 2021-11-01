package entities

type SignUpResponse struct {
	Kode    uint   `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
