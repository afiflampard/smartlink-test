package entities

type Layanan struct {
	Nama  string `json:"nama"`
	Unit  string `json:"unit"`
	Harga string `json:"harga"`
}

type LayananResponse struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Data   DataResponse `json:"data"`
}

type DataResponse struct {
	ID     string `json:"id"`
	Nama   string `json:"nama"`
	Unit   string `json:"unit"`
	Harga  string `json:"harga"`
	UserID string `json:"User_id"`
}
