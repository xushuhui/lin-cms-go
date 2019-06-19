package request

type CreateBook struct {
}
type GetBook struct {
}
type UpdateBook struct {
	Id      int    `json:"id"valid:"Required;Numeric"`
	Title   string `json:"title"valid:"Required"`
	Author  string `json:"author"valid:"Required"`
	Summary string `json:"summary"valid:"Required"`
	Image   string `json:"image"valid:"Required"`
}
