package service

type LogPage struct {
	Total int        `json:"total"`
	Items []LogItems `json:"items"`
	Page  int        `json:"page"`
	Count int        `json:"count"`
}
type LogItems struct {
	ID         int    `json:"id"`
	Message    string `json:"message"`
	UserID     int    `json:"user_id"`
	Username   string `json:"username"`
	StatusCode int    `json:"status_code"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	Permission string `json:"permission"`
}
type UserPage struct {
	Total int         `json:"total"`
	Items []UserItems `json:"items"`
	Page  int         `json:"page"`
	Count int         `json:"count"`
}
type Group struct {
	ID   int    `json:"id"`
	Info string `json:"info"`
	Name string `json:"name"`
}
type UserItems struct {
	ID       int         `json:"id"`
	Username string      `json:"username"`
	Nickname string      `json:"nickname"`
	Avatar   interface{} `json:"avatar"`
	Email    interface{} `json:"email"`
	Group    []Group     `json:"group"`
}
