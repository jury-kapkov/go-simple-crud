package todo

type TodoList struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ItemsList struct {
	Id     int
	ListId int
	ItemId int
}
