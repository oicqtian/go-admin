package request

type IdRequest struct {
	ID  int64   `json:"id"`
	IDS []int64 `json:"ids"`
}

type PageRequest struct {
	Page     int `json:page`
	PageSize int `json:pageSize`
}
