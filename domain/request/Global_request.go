package request

type DataTableParam struct {
	LastID   string `json:"last_id"`
	PerPage  int    `json:"per_page"`
	Page     int    `json:"page"`
	OrderBy  string `json:"order_by"`
	OrderDir string `json:"order_dir"`
}
