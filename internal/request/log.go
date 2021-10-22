package request

type GetLogs struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Name  string `json:"name"`

	Pages
}
type SearchLogs struct {
	Keyword string `json:"keyword"`
	GetLogs
}
type GetLogUsers struct {
	Pages
}
