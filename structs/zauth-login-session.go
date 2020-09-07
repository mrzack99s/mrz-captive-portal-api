package structs

type ZAuthLoginSession struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ip_address" db:"telcode" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

func (b *ZAuthLoginSession) TableName() string {
	return "loginsession"
}
