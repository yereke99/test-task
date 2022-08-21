package dto

type DataRequest struct {
	Method string `json:"method" binding:"required"`
	Url    string `json:"url" binding:"required"`
	Data   string `json:"data" binding:"required"`
}

type DataResponse struct {
	Status  string `json:"status" binding:"required"`
	Headers string `json:"headers" binding:"required"`
	Result  string `json:"result" binding:"required"`
}
