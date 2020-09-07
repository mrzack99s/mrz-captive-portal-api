package structs

type ZAuthCheck struct {
	Password  string `json:"Password" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	IPAddress string `json:"IPAddress" binding:"required"`
}
