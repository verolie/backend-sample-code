package modelResponse

type DetailGetResponse struct {
	Detail    interface{} `json:"detail"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
}