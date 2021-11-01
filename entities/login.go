package entities

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ID       string `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
