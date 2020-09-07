package structs

type ZAuthEAP struct {
	Username  string `json:"Username" binding:"required"`
	IPAddress string `json:"IPAddress" binding:"required"`
}
