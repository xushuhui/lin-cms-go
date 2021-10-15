package request

type UpdateBook struct {
	Title   string `json:"title" validate:"required" comment:"书籍标题"`
	Author  string `json:"author" validate:"required" comment:"书籍作者"`
	Summary string `json:"summary" validate:"required" comment:"书籍简介"`
	Image   string `json:"image" validate:"required" comment:"书籍封面图"`
}

type CreateBook struct {
	Title   string `json:"title" validate:"required" comment:"书籍标题"`
	Author  string `json:"author" validate:"required" comment:"书籍作者"`
	Summary string `json:"summary" validate:"required" comment:"书籍简介"`
	Image   string `json:"image" validate:"required" comment:"书籍封面图"`
}
type GetBooks struct {
	Size int `json:"size" binding:"required" comment:"数量"`
	Page int `json:"page" binding:"required" comment:"页码"`
}
