package transfers

type PaginateData struct {
	Data  []TransactionJson `json:"data"`
	Total int               `json:"total"`
	Page  int               `json:"page"`
}
