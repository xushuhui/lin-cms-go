package request

type UpdateBook struct {
	Id      int    `json:"id" validate:"required" comment:"书籍id"`
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

type DeleteBook struct {
	Id int `json:"id" validate:"required" comment:"书籍id"`
}
