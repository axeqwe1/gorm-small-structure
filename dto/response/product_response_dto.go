package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // `omitempty` ใช้เพื่อละเว้น field นี้หากไม่มีค่า
}
