package structs

type ZAuthIPAddress struct {
	IPAddress string `json:"IPAddress" binding:"required"`
}
